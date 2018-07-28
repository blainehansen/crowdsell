package main

import (
	"fmt"
	"bytes"

	"github.com/gin-gonic/gin"
	"github.com/badoux/checkmail"
	"gopkg.in/doug-martin/goqu.v4"
)

type SignedUser struct {
	Name string
	UrlSlug string
	Email string
	Token string
}

var _ r = route(POST, "/create-user", func(c *gin.Context) {
	newUser := struct {
		Name string `binding:"required"`
		Email string `binding:"required"`
		Password []byte `binding:"required"`
	}{}
	if err := c.ShouldBindJSON(&newUser); err != nil {
		c.AbortWithError(422, err); return
	}

	if err := checkmail.ValidateFormat(newUser.Email); err != nil {
		c.AbortWithError(400, err); return
	}

	hashedPassword, hashError := hashPassword(&newUser.Password)
	if hashError != nil {
		c.AbortWithError(500, hashError); return
	}

	var userSlug string
	found, err := db.From("users").Returning(goqu.I("url_slug")).Insert(
		goqu.Record{ "name": newUser.Name, "email": newUser.Email, "password": hashedPassword },
	).ScanVal(&userSlug)
	if err != nil || !found {
		c.AbortWithError(500, err); return
	}

	authTokenString, issueError := issueAuthToken(userSlug)
	if issueError != nil {
		c.AbortWithError(500, issueError); return
	}

	c.JSON(200, SignedUser { Name: newUser.Name, Email: newUser.Email, UrlSlug: userSlug, Token: authTokenString })
})



var _ r = route(POST, "/login", func(c *gin.Context) {
	loginAttempt := struct {
		Email string `binding:"required"`
		Password []byte `binding:"required"`
	}{}
	if err := c.ShouldBindJSON(&loginAttempt); err != nil {
		c.AbortWithError(422, err); return
	}

	databaseUser := struct {
		Slug string
		UrlSlug string
		Name string
		Email string
		Password []byte
	}{}
	userFound, queryError := db.From("users").Where(goqu.Ex{ "email": loginAttempt.Email }).ScanStruct(&databaseUser)
	if queryError != nil {
		c.AbortWithError(500, queryError); return
	}
	if !userFound {
		c.AbortWithStatus(403); return
	}

	if !verifyPassword(&databaseUser.Password, &loginAttempt.Password) {
		c.AbortWithStatus(403); return
	}

	authTokenString, issueError := issueAuthToken(databaseUser.Slug)
	if issueError != nil {
		c.AbortWithError(500, issueError); return
	}
	c.JSON(200, SignedUser { Name: databaseUser.Name, UrlSlug: databaseUser.UrlSlug, Email: loginAttempt.Email, Token: authTokenString })
})

var _ r = authRoute(POST, "/users/change-slug", func(c *gin.Context) {
	userId := c.MustGet("userId").(int64)

	slugPayload := struct {
		UrlSlug string `binding:"required"`
	}{}
	if err := c.ShouldBindJSON(&slugPayload); err != nil {
		c.AbortWithError(422, err); return
	}

	result, err := db.From("users").Where(
		goqu.Ex{ "id": userId },
	).Update(
		goqu.Record{ "url_slug": slugPayload.UrlSlug },
	).Exec()
	if err != nil {
		c.AbortWithError(500, err); return
	}
	if rowsAffected, err := result.RowsAffected(); rowsAffected == 0 || err != nil {
		c.AbortWithStatus(403); return
	}

	c.Status(204)
})


var _ r = route(POST, "/users/forgot-password", func(c *gin.Context) {
	forgottenEmailPayload := struct {
		Email string `binding:"required"`
	}{}
	if err := c.ShouldBindJSON(&forgottenEmailPayload); err != nil {
		c.AbortWithError(422, err); return
	}

	forgottenEmail := forgottenEmailPayload.Email
	if err := checkmail.ValidateFormat(forgottenEmail); err != nil {
		c.AbortWithError(400, err); return
	}

	forgotPasswordToken, generationError := generateRandomToken()
	if generationError != nil {
		c.AbortWithError(500, generationError)
	}

	result, err := db.From("users").Where(
		goqu.Ex{ "email": forgottenEmail },
	).Update(
		goqu.Record{ "forgot_password_token": forgotPasswordToken },
	).Exec()
	if err != nil {
		c.AbortWithError(500, err); return
	}
	if rowsAffected, err := result.RowsAffected(); rowsAffected == 0 || err != nil {
		// should we pretend like everything's okay? this would essentially pretend like there's a user with that email, and not provide any information about the user base
		c.Status(204); return
	}


	recoveryToken := []byte(fmt.Sprintf("%s:%s", forgottenEmail, forgotPasswordToken))
	message := fmt.Sprintf(`%s%s/recover-password?t=%s`, environment["SERVER_PROTOCOL"], environment["SERVER_DOMAIN"], encodeBase64(recoveryToken))
	if err := sendMessage("no-reply@crowdsell.io", "Forgot Password", message, forgottenEmail); err != nil {
		c.AbortWithError(500, err)
	}

	c.Status(204)
})

var _ r = route(POST, "/users/recover-password", func(c *gin.Context) {
	recoveryTokenPayload := struct {
		RecoveryToken []byte `binding:"required"`
		NewPassword []byte `binding:"required"`
	}{}
	if err := c.ShouldBindJSON(&recoveryTokenPayload); err != nil {
		c.AbortWithError(422, err); return
	}

	recoveryToken, err := decodeBase64(recoveryTokenPayload.RecoveryToken)
	if err != nil {
		c.AbortWithStatus(400); return
	}

	lastIndex := bytes.LastIndexByte(recoveryToken, ':')
	if lastIndex < 1 {
		c.AbortWithStatus(400); return
	}

	forgottenEmail := string(recoveryToken[:lastIndex])
	forgotPasswordToken := recoveryToken[lastIndex + 1 :]


	hashedPassword, hashError := hashPassword(&recoveryTokenPayload.NewPassword)
	if hashError != nil {
		c.AbortWithError(500, hashError); return
	}


	databaseUser := struct {
		Slug string
		UrlSlug string
		Name string
		Email string
	}{}
	found, err := db.From("users").Where(
		goqu.Ex{ "email": forgottenEmail, "forgot_password_token": forgotPasswordToken },
	).Returning("slug", "url_slug", "name", "email").Update(
		goqu.Record{ "forgot_password_token": nil, "password": hashedPassword },
	).ScanStruct(&databaseUser)
	if err != nil {
		c.AbortWithError(500, err); return
	}
	if !found {
		c.AbortWithStatus(403); return
	}

	authTokenString, issueError := issueAuthToken(databaseUser.Slug)
	if issueError != nil {
		c.AbortWithError(500, issueError); return
	}

	c.JSON(200, SignedUser { Name: databaseUser.Name, UrlSlug: databaseUser.UrlSlug, Email: databaseUser.Email, Token: authTokenString })
})

func VerifyTokenMiddleWare(c *gin.Context) {
	authHeader := c.GetHeader("Authorization")

	userId, userSlug, verifyError := verifyAuthToken(authHeader)
	switch verifyError {
		case nil:
			break
		case InvalidTokenError:
			c.AbortWithStatus(400); return
		case ExpiredTokenError:
			c.AbortWithStatus(401); return
		case UnauthorizedError:
			c.AbortWithStatus(403); return
		default:
			c.AbortWithError(500, verifyError); return
	}

	c.Set("userId", userId)
	c.Set("userSlug", userSlug)
	// c.Set("userRole", userRole)
	c.Next()
}
