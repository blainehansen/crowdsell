package main

import (
	"fmt"
	"time"
	"strings"
	"errors"

	"github.com/gin-gonic/gin"

	"encoding/base64"
	"github.com/json-iterator/go"
	"github.com/lhecker/argon2"

	"crypto"
	_ "crypto/sha256"
	"crypto/hmac"
)


type SignedUser struct {
	Name string `json:"name"`
	Email string `json:"email"`
	Token string `json:"token"`
}
type NewUserJson struct {
	Name string `json:"name" binding:"required"`
	Email string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}
var _ r = route(POST, "/create-user", func(c *gin.Context) {
	var newUser NewUserJson
	if err := c.ShouldBindJSON(&newUser); err != nil {
		c.AbortWithError(422, err)
		return
	}

	hashedPassword, hashError := hashPassword([]byte(newUser.Password))
	if hashError != nil {
		c.AbortWithError(500, hashError)
		return
	}

	user := User { Name: newUser.Name, Email: newUser.Email, Password: string(hashedPassword) }
	if err := db.Create(&user).Error; err != nil {
		c.AbortWithError(500, err)
		return
	}

	fmt.Println("user.Id", user.Id)
	// TODO don't have an empty role
	authTokenString, issueError := issueAuthToken(user.Id, "")
	if issueError != nil {
		c.AbortWithError(500, issueError)
		return
	}

	c.JSON(200, SignedUser { Name: user.Name, Email: user.Email, Token: authTokenString })
})



type LoginJson struct {
	Email string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}
type LoginUser struct {
	Id uint32
	Password []byte
}
var _ r = route(POST, "/login", func(c *gin.Context) {
	var loginAttempt LoginJson
	if err := c.ShouldBindJSON(&loginAttempt); err != nil {
		c.AbortWithError(422, err)
		return
	}

	var user LoginUser
	queryResult := db.Table("users").Select("id, password").Where("email = ?", loginAttempt.Email).Scan(&user)
	if queryResult.RecordNotFound() {
		c.AbortWithStatus(403)
		return
	}
	if queryResult.Error != nil {
		c.AbortWithError(500, queryResult.Error)
		return
	}

	if !verifyPassword(user.Password, []byte(loginAttempt.Password)) {
		c.AbortWithStatus(403)
		return
	}

	// TODO don't have an empty role
	authTokenString, issueError := issueAuthToken(user.Id, "")
	if issueError != nil {
		c.AbortWithError(500, issueError)
		return
	}
	c.JSON(200, authTokenString)
})

var _ r = authRoute(GET, "/special", func(c *gin.Context) {
	userId := c.MustGet("userId").(uint32)
	fmt.Println(userId)
	c.JSON(200, "stuff")
})


func generateKey() *[]byte {
	key := []byte("somepass")
	return &key
}
var privateKey *[]byte = generateKey()


type AuthToken struct {
	I uint32
	E int64
	R string
}
func (tok *AuthToken) Id() uint32 {
	return tok.I
}
func (tok *AuthToken) Expires() int64 {
	return tok.E
}
func (tok *AuthToken) Role() string {
	return tok.R
}

func issueAuthToken(userId uint32, role string) (string, error) {
	// get tomorrow unix time
	tomorrow := time.Now().Add(time.Duration(24) * time.Hour).Unix()

	// create an authtoken
	token := AuthToken{ userId, tomorrow, role }

	// json serialize it and base64 encode it
	serializedToken, serializationError := jsoniter.Marshal(token)
	if serializationError != nil {
		return "", serializationError
	}
	fmt.Printf("%s\n", serializedToken)

	encodedToken := encodeBase64(serializedToken)

	// create a signature of it
	signer := hmac.New(crypto.SHA256.New, *privateKey)
	signer.Write(encodedToken)
	signature := signer.Sum(nil)

	// base64 encode the signature
	encodedSignature := encodeBase64(signature)

	// put the two together
	return fmt.Sprintf("%s.%s", encodedToken, encodedSignature), nil
}

var InvalidTokenError error = errors.New("InvalidTokenError")
var UnauthorizedError error = errors.New("UnauthorizedError")
var ExpiredTokenError error = errors.New("ExpiredTokenError")

func VerifyTokenMiddleWare(c *gin.Context) {
	authHeader := c.GetHeader("Authorization")

	userId, userRole, verifyError := verifyAuthToken(authHeader)
	switch verifyError {
		case nil:
			break
		case InvalidTokenError:
			c.AbortWithStatus(400)
			return
		case ExpiredTokenError:
			c.AbortWithStatus(401)
			return
		case UnauthorizedError:
			c.AbortWithStatus(403)
			return
		default:
			c.AbortWithError(500, verifyError)
			return
	}

	c.Set("userId", userId)
	c.Set("userRole", userRole)
	c.Next()
}

func verifyAuthToken(token string) (uint32, string, error) {
	// p := fmt.Println
	// f := fmt.Printf

	// split the segments
	segments := strings.SplitN(token, ".", 2)
	// p(segments)

	if len(segments) != 2 {
		return 0, "", InvalidTokenError
	}
	proposedEncodedToken := []byte(segments[0])
	proposedEncodedSignature := []byte(segments[1])
	// f("proposedEncodedToken %s\n", proposedEncodedToken)
	// f("proposedEncodedSignature %s\n", proposedEncodedSignature)

	// decode the signature
	proposedSignature, signatureDecodeError := decodeBase64(proposedEncodedSignature)
	if signatureDecodeError != nil {
		return 0, "", InvalidTokenError
	}
	// f("proposedSignature %s\n", proposedSignature)
	// p("proposedSignature", proposedSignature)

	signer := hmac.New(crypto.SHA256.New, *privateKey)
	// sign the incoming encoded token
	signer.Write(proposedEncodedToken)
	actualSignature := signer.Sum(nil)

	// f("actualSignature %s\n", actualSignature)
	// p("actualSignature", actualSignature)
	// and see if the signature they have matches what we would produce
	if !hmac.Equal(proposedSignature, actualSignature) {
		return 0, "", UnauthorizedError
	}

	proposedSerializedToken, tokenDecodeError := decodeBase64(proposedEncodedToken)
	if tokenDecodeError != nil {
		return 0, "", InvalidTokenError
	}
	// f("proposedSerializedToken %s\n", proposedSerializedToken)

	var successfulToken AuthToken
	if deserializationError := jsoniter.Unmarshal(proposedSerializedToken, &successfulToken); deserializationError != nil {
		return 0, "", InvalidTokenError
	}
	// p("successfulToken", successfulToken)

	// check the expires
	if successfulToken.Expires() <= time.Now().Unix() {
		return 0, "", ExpiredTokenError
	}

	return successfulToken.Id(), successfulToken.Role(), nil
}


var passwordHasher argon2.Config = argon2.DefaultConfig()

func hashPassword(password []byte) ([]byte, error) {
	encodedPassword, encodingError := passwordHasher.HashEncoded(password)
	argon2.SecureZeroMemory(password)
	return encodedPassword, encodingError
}

func verifyPassword(encodedPassword []byte, trialPassword []byte) bool {
	match, err := argon2.VerifyEncoded(trialPassword, encodedPassword)
	if err != nil {
		return false
	}
	return match
}
var localEncoding *base64.Encoding = base64.RawURLEncoding

func encodeBase64(data []byte) []byte {
	encodedData := make([]byte, localEncoding.EncodedLen(len(data)))
	localEncoding.Encode(encodedData, data)
	return encodedData
}

func decodeBase64(data []byte) ([]byte, error) {
	decodedData := make([]byte, localEncoding.DecodedLen(len(data)))
	_, decodeError := localEncoding.Decode(decodedData, data)
	return decodedData, decodeError
}