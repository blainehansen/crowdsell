package main

import (
	"fmt"
	"bytes"

	"github.com/gin-gonic/gin"
	"github.com/badoux/checkmail"
	// "gopkg.in/doug-martin/goqu.v4"
	"github.com/blainehansen/goqu"
)

type SignedUser struct {
	Name string
	Slug string
	Email string
	Token string
}

var _ r = route(POST, "/create-user", func(c *gin.Context) {
	newUser := struct {
		Name string `binding:"required"`
		Email string `binding:"required"`
		Password string `binding:"required"`
	}{}
	if err := c.ShouldBindJSON(&newUser); err != nil {
		c.AbortWithError(422, err); return
	}

	if err := checkmail.ValidateFormat(newUser.Email); err != nil {
		c.AbortWithError(400, err); return
	}

	newUserPassword := []byte(newUser.Password)
	hashedPassword, hashError := hashPassword(&newUserPassword)
	if hashError != nil {
		c.AbortWithError(500, hashError); return
	}

	var userId string
	found, err := db.From("person").Returning(goqu.I("id")).Insert(
		goqu.Record{ "name": newUser.Name, "email": newUser.Email, "password": hashedPassword },
	).ScanVal(&userId)
	if err != nil {
		c.AbortWithError(500, err); return
	}
	if !found {
		c.AbortWithError(500, fmt.Errorf("userId not found? %s", userId)); return
	}

	authTokenString, issueError := issueAuthToken(userId)
	if issueError != nil {
		c.AbortWithError(500, issueError); return
	}

	c.JSON(200, SignedUser { Name: newUser.Name, Email: newUser.Email, Slug: userId, Token: authTokenString })
})


var _ r = route(POST, "/login", func(c *gin.Context) {
	loginAttempt := struct {
		Email string `binding:"required"`
		Password string `binding:"required"`
	}{}
	if err := c.ShouldBindJSON(&loginAttempt); err != nil {
		c.AbortWithError(422, err); return
	}

	loginAttemptPassword := []byte(loginAttempt.Password)
	valid, databaseUser := checkUserPassword(c, goqu.I("email").Eq(loginAttempt.Email), &loginAttemptPassword)
	if !valid { return }

	authTokenString, issueError := issueAuthToken(databaseUser.Slug)
	if issueError != nil {
		c.AbortWithError(500, issueError); return
	}

	c.JSON(200, SignedUser { Name: databaseUser.Name, Slug: databaseUser.Slug, Email: loginAttempt.Email, Token: authTokenString })
})


var _ r = authRoute(PUT, "/user/password", func(c *gin.Context) {
	userId := c.MustGet("userId").(string)

	passwordPayload := struct {
		OldPassword string `binding:"required"`
		NewPassword string `binding:"required"`
	}{}
	if err := c.ShouldBindJSON(&passwordPayload); err != nil {
		c.AbortWithError(422, err); return
	}


	oldPassword := []byte(passwordPayload.OldPassword)
	valid, databaseUser := checkUserPassword(c, goqu.I("id").Eq(userId), &oldPassword)
	if !valid { return }

	// hash the new password
	newPassword := []byte(passwordPayload.NewPassword)
	hashedPassword, hashError := hashPassword(&newPassword)
	if hashError != nil {
		c.AbortWithError(500, hashError); return
	}

	updateQuery := db.From("person").Where(
		goqu.I("id").Eq(userId),
	).Update(
		goqu.Record{ "password": hashedPassword },
	)

	if !doExec(c, updateQuery) { return }

	// give them a new signed user
	authTokenString, issueError := issueAuthToken(databaseUser.Slug)
	if issueError != nil {
		c.AbortWithError(500, issueError); return
	}

	c.JSON(200, SignedUser { Name: databaseUser.Name, Slug: databaseUser.Slug, Email: databaseUser.Email, Token: authTokenString })
})


type readyToSignUser struct {
	Slug string
	Name string
	Email string
	Password []byte
}

func checkUserPassword(c *gin.Context, condition goqu.BooleanExpression, attemptPassword *[]byte) (bool, readyToSignUser) {
	databaseUser := readyToSignUser{}
	userFound, queryError := db.From("person").Where(condition).ScanStruct(&databaseUser)
	if queryError != nil {
		c.AbortWithError(500, queryError);
		return false, databaseUser
	}
	if !userFound {
		c.AbortWithStatus(403);
		return false, databaseUser
	}

	if !verifyPassword(&databaseUser.Password, attemptPassword) {
		c.AbortWithStatus(403);
		return false, databaseUser
	}

	return true, databaseUser
}



var _ r = route(POST, "/user/password/forgot", func(c *gin.Context) {
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

	result, err := db.From("person").Where(
		goqu.I("email").Eq(forgottenEmail),
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

var _ r = route(POST, "/user/password/recover", func(c *gin.Context) {
	recoveryTokenPayload := struct {
		RecoveryToken string `binding:"required"`
		NewPassword string `binding:"required"`
	}{}
	if err := c.ShouldBindJSON(&recoveryTokenPayload); err != nil {
		c.AbortWithError(422, err); return
	}

	recoveryToken, err := decodeBase64([]byte(recoveryTokenPayload.RecoveryToken))
	if err != nil {
		c.AbortWithStatus(400); return
	}

	lastIndex := bytes.LastIndexByte(recoveryToken, ':')
	if lastIndex < 1 {
		c.AbortWithStatus(400); return
	}

	forgottenEmail := string(recoveryToken[:lastIndex])
	forgotPasswordToken := recoveryToken[lastIndex + 1 :]

	recoveryNewPassword := []byte(recoveryTokenPayload.NewPassword)
	hashedPassword, hashError := hashPassword(&recoveryNewPassword)
	if hashError != nil {
		c.AbortWithError(500, hashError); return
	}

	databaseUser := struct {
		Id string
		Slug string
		Name string
		Email string
	}{}

	found, err := db.From("person").Where(
		goqu.I("email").Eq(forgottenEmail), goqu.I("forgot_password_token").Eq(forgotPasswordToken),
	).Returning(goqu.I("id"), goqu.I("slug"), goqu.I("name"), goqu.I("email")).Update(
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

	c.JSON(200, SignedUser { Name: databaseUser.Name, Slug: databaseUser.Slug, Email: databaseUser.Email, Token: authTokenString })
})


func VerifyTokenMiddleWare(c *gin.Context) {
	authHeader := c.GetHeader("Authorization")

	userId, verifyError := verifyAuthToken(authHeader)
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
	// c.Set("userRole", userRole)
	c.Next()
}
