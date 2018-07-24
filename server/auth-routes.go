package main

import (
	// "fmt"
	// "bytes"

	"github.com/gin-gonic/gin"
	// "github.com/badoux/checkmail"
)

type SignedUser struct {
	Name string
	Slug string
	Email string
	Token string
}

var _ r = route(POST, "/create-user", func(c *gin.Context) {
	newUser := struct {
		Name *string `binding:"required"`
		Email string `binding:"required"`
		Password []byte `binding:"required"`
	}{}
	if err := c.ShouldBindJSON(&newUser); err != nil {
		c.AbortWithError(422, err); return
	}

	hashedPassword, hashError := hashPassword(&newUser.Password)
	if hashError != nil {
		c.AbortWithError(500, hashError); return
	}

	createdUser := User { Name: newUser.Name, Email: newUser.Email, Password: hashedPassword }
	if err := dbUserStore.Insert(&createdUser); err != nil {
		c.AbortWithError(500, err); return
	}

	userInternalSlug, authTokenString, issueError := issueAuthTokenForId(createdUser.Id)
	if issueError != nil {
		c.AbortWithError(500, issueError); return
	}

	c.JSON(200, SignedUser { Name: *createdUser.Name, Slug: userInternalSlug, Email: createdUser.Email, Token: authTokenString })
})



var _ r = route(POST, "/login", func(c *gin.Context) {
	loginAttempt := struct {
		Email string `binding:"required"`
		Password []byte `binding:"required"`
	}{}
	if err := c.ShouldBindJSON(&loginAttempt); err != nil {
		c.AbortWithError(422, err); return
	}

	databaseUser, queryError := dbUserStore.FindOne(
		NewUserQuery().Select(
			Schema.User.Slug, Schema.User.InternalSlug, Schema.User.Name, Schema.User.Email, Schema.User.Password,
		).FindByEmail(loginAttempt.Email),
	)
	if queryError != nil {
		c.AbortWithError(500, queryError); return
	}
	if databaseUser == nil {
		c.AbortWithStatus(403); return
	}

	if !verifyPassword(&databaseUser.Password, &loginAttempt.Password) {
		c.AbortWithStatus(403); return
	}

	authTokenString, issueError := issueAuthToken(databaseUser.InternalSlug)
	if issueError != nil {
		c.AbortWithError(500, issueError); return
	}
	c.JSON(200, SignedUser { Name: *databaseUser.Name, Slug: databaseUser.Slug, Email: loginAttempt.Email, Token: authTokenString })
})

var _ r = authRoute(POST, "/users/change-slug", func(c *gin.Context) {
	userId := c.MustGet("userId").(int64)

	slugPayload := struct {
		Slug string `binding:"required"`
	}{}
	if err := c.ShouldBindJSON(&slugPayload); err != nil {
		c.AbortWithError(422, err); return
	}

	user := User { Id: userId, Slug: slugPayload.Slug }
	rowsUpdated, err := dbUserStore.Update(&user, Schema.User.Slug)
	if err != nil {
		c.AbortWithError(500, err); return
	}
	if rowsUpdated == 0 {
		c.AbortWithStatus(403); return
	}

	c.Status(204)
})


// var _ r = route(POST, "/users/forgot-password", func(c *gin.Context) {
// 	forgottenEmailPayload := struct {
// 		Email string `binding:"required"`
// 	}{}
// 	if err := c.ShouldBindJSON(&forgottenEmailPayload); err != nil {
// 		c.AbortWithError(422, err); return
// 	}

// 	forgottenEmail := forgottenEmailPayload.Email
// 	if err := checkmail.ValidateFormat(forgottenEmail); err != nil {
// 		c.AbortWithError(400, err)
// 	}

// 	user, err := dbUserStore.FindOne(
// 		NewUserQuery().FindByEmail(forgottenEmail).Select(Schema.User.Id, Schema.User.ForgotPasswordToken),
// 	)
// 	if err != nil {
// 		c.AbortWithError(500, err); return
// 	}
// 	if user == nil {
// 		c.Status(204); return
// 	}

// 	forgotPasswordToken, generationError := generateRandomToken()
// 	if generationError != nil {
// 		c.AbortWithError(500, generationError)
// 	}
// 	user.ForgotPasswordToken = &forgotPasswordToken

// 	rowsUpdated, err := dbUserStore.Update(user, Schema.User.ForgotPasswordToken)
// 	if err != nil || rowsUpdated == 0 {
// 		c.AbortWithStatus(500); return
// 	}

// 	recoveryToken := []byte(fmt.Sprintf("%s:%s", forgottenEmail, forgotPasswordToken))
// 	message := fmt.Sprintf(`%s%s/recover-password?t=%s`, environment["SERVER_PROTOCOL"], environment["SERVER_DOMAIN"], encodeBase64(recoveryToken))

// 	// TODO this will be real
// 	// if err := sendMessage("no-reply@crowdsell.io", "Forgot Password", message, forgottenEmail); err != nil {
// 	// 	c.AbortWithError(500, err)
// 	// }

// 	c.Status(204)
// })

// var _ r = route(POST, "/users/recover-password", func(c *gin.Context) {
// 	recoveryTokenPayload := struct {
// 		RecoveryToken *[]byte `binding:"required"`
// 		NewPassword *[]byte `binding:"required"`
// 	}{}
// 	if err := c.ShouldBindJSON(&recoveryTokenPayload); err != nil {
// 		c.AbortWithError(422, err); return
// 	}

// 	recoveryToken, err := decodeBase64(*recoveryTokenPayload.RecoveryToken)
// 	if err != nil {
// 		c.AbortWithStatus(400); return
// 	}

// 	lastIndex := bytes.LastIndexByte(recoveryToken, ':')
// 	if lastIndex < 1 {
// 		c.AbortWithStatus(400); return
// 	}

// 	forgottenEmail := string(recoveryToken[:lastIndex])
// 	forgotPasswordToken := recoveryToken[lastIndex:]

// 	user, err := dbUserStore.FindOne(
// 		NewUserQuery().FindByEmail(forgottenEmail).Select(Schema.User.Id, Schema.User.Name, Schema.User.Slug, Schema.User.InternalSlug, Schema.User.Password, Schema.User.ForgotPasswordToken),
// 	)
// 	if err != nil || user == nil {
// 		c.AbortWithStatus(400); return
// 	}

// 	if !bytes.Equal(*user.ForgotPasswordToken, forgotPasswordToken) {
// 		c.AbortWithStatus(400); return
// 	}


// 	hashedPassword, hashError := hashPassword(recoveryTokenPayload.NewPassword)
// 	if hashError != nil {
// 		c.AbortWithError(500, hashError); return
// 	}

// 	// TODO or something
// 	user.ForgotPasswordToken = nil
// 	user.Password = hashedPassword

// 	rowsUpdated, err := dbUserStore.Update(user, Schema.User.Password, Schema.User.ForgotPasswordToken)
// 	if rowsUpdated == 0 || err != nil {
// 		c.AbortWithStatus(500); return
// 	}

// 	authTokenString, issueError := issueAuthToken(user.InternalSlug)
// 	if issueError != nil {
// 		c.AbortWithError(500, issueError); return
// 	}

// 	c.JSON(200, SignedUser { Name: *user.Name, Slug: user.Slug, Email: user.Email, Token: authTokenString })
// })

func VerifyTokenMiddleWare(c *gin.Context) {
	authHeader := c.GetHeader("Authorization")

	userId, userInternalSlug, verifyError := verifyAuthToken(authHeader)
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
	c.Set("userInternalSlug", userInternalSlug)
	// c.Set("userRole", userRole)
	c.Next()
}
