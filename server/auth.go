package main

import (
	"fmt"
	"time"
	"strings"
	"errors"

	"github.com/gin-gonic/gin"

	"encoding/base64"
	"github.com/lhecker/argon2"
	"github.com/json-iterator/go"

	"github.com/speps/go-hashids"

	"crypto"
	_ "crypto/sha256"
	"crypto/hmac"
)


type SignedUser struct {
	Name string `json:"name"`
	Slug string `json:"slug"`
	Email string `json:"email"`
	Token string `json:"token"`
}

var _ r = route(POST, "/create-user", func(c *gin.Context) {
	newUser := struct {
		Name string `json:"name" binding:"required"`
		Email string `json:"email" binding:"required"`
		Password []byte `json:"password" binding:"required"`
	}{}
	if err := c.ShouldBindJSON(&newUser); err != nil {
		c.AbortWithError(422, err); return
	}

	hashedPassword, hashError := hashPassword(&newUser.Password)
	if hashError != nil {
		c.AbortWithError(500, hashError); return
	}

	createdUser := User { Name: newUser.Name, Email: newUser.Email, Password: hashedPassword }
	if err := db.Create(&createdUser).Error; err != nil {
		c.AbortWithError(500, err); return
	}
	fmt.Println(createdUser.Id)
	fmt.Println(createdUser.Slug)

	authTokenString, issueError := issueAuthToken(createdUser.Id)
	if issueError != nil {
		c.AbortWithError(500, issueError); return
	}

	c.JSON(200, SignedUser { Name: createdUser.Name, Slug: createdUser.Slug, Email: createdUser.Email, Token: authTokenString })
})



var _ r = route(POST, "/login", func(c *gin.Context) {
	loginAttempt := struct {
		Email string `json:"email" binding:"required"`
		Password []byte `json:"password" binding:"required"`
	}{}
	if err := c.ShouldBindJSON(&loginAttempt); err != nil {
		c.AbortWithError(422, err); return
	}

	databaseUser := struct {
		Id int64
		Slug string
		Name string
		Password []byte
	}{}
	queryResult := db.Table("users").Select("id, slug, name, password").Where("email = ?", &loginAttempt.Email).Scan(&databaseUser)
	if queryResult.RecordNotFound() {
		c.AbortWithStatus(403); return
	}
	if queryResult.Error != nil {
		c.AbortWithError(500, queryResult.Error); return
	}

	if !verifyPassword(&databaseUser.Password, &loginAttempt.Password) {
		c.AbortWithStatus(403); return
	}

	authTokenString, issueError := issueAuthToken(databaseUser.Id)
	if issueError != nil {
		c.AbortWithError(500, issueError); return
	}
	c.JSON(200, SignedUser { Name: databaseUser.Name, Slug: databaseUser.Slug, Email: loginAttempt.Email, Token: authTokenString })
})

var _ r = authRoute(POST, "/users/:userSlug/slug", func(c *gin.Context) {
	userId := c.MustGet("userId").(int64)

	slugPayload := struct {
		Slug string `json:"slug" binding:"required"`
	}{}
	if err := c.ShouldBindJSON(&slugPayload); err != nil {
		c.AbortWithError(422, err); return
	}

	databaseUser := struct {
		Id int64
		Slug string
		Name string
		Email string
	}{}
	if db.Table("users").Select("id, slug, name, email").Where("id = ?", userId).Scan(&databaseUser).Error != nil {
		c.AbortWithStatus(500); return
	}

	updateUser := User{}
	updateUser.Id = databaseUser.Id
	db.Model(&updateUser).Update("slug", slugPayload.Slug)
	fmt.Println(databaseUser.Id)
	fmt.Println(databaseUser.Slug)

	authTokenString, issueError := issueAuthToken(databaseUser.Id)
	if issueError != nil {
		c.AbortWithError(500, issueError); return
	}
	c.JSON(200, SignedUser { Name: databaseUser.Name, Slug: databaseUser.Slug, Email: databaseUser.Email, Token: authTokenString })
})



// TODO populate both of these with environment or config variables instead somehow
var privateKey *[]byte = func() *[]byte {
	key := []byte("somepass")
	return &key
}()

var HashIDData *hashids.HashIDData = func() *hashids.HashIDData {
	internalHashIdData := hashids.HashIDData {
		Alphabet: "abcdefghijklmnopqrstuvwxyz-ABCDEFGHIJKLMNOPQRSTUVWXYZ",
		MinLength: 8,
		Salt: "id& obfuscation sys$tem here",
	}
	return &internalHashIdData
}()


type AuthToken struct {
	I string
	E int64
	// R map[int]int
}
func (tok *AuthToken) InternalSlug() string {
	return tok.I
}
func (tok *AuthToken) Expires() int64 {
	return tok.E
}
// func (tok *AuthToken) Role() string {
// 	return tok.R
// }

func issueAuthToken(userId int64) (string, error) {
	// get tomorrow unix time
	tomorrow := time.Now().Add(time.Duration(24) * time.Hour).Unix()

	hashId, hashIdError := hashids.NewWithData(HashIDData)
	userInternalSlug, encodeError := hashId.EncodeInt64([]int64{userId})
	if hashIdError != nil || encodeError != nil {
		return "", encodeError
	}

	// create an authtoken
	token := AuthToken{ userInternalSlug, tomorrow }

	// json serialize it and base64 encode it
	serializedToken, serializationError := jsoniter.Marshal(token)
	if serializationError != nil {
		return "", serializationError
	}

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

	userId, userInternalSlug, verifyError := verifyAuthToken(authHeader)
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
	c.Set("userInternalSlug", userInternalSlug)
	// c.Set("userRole", userRole)
	c.Next()
}

func verifyAuthToken(token string) (int64, string, error) {
	// split the segments
	segments := strings.SplitN(token, ".", 2)

	if len(segments) != 2 {
		return 0, "", InvalidTokenError
	}
	proposedEncodedToken := []byte(segments[0])
	proposedEncodedSignature := []byte(segments[1])

	// decode the signature
	proposedSignature, signatureDecodeError := decodeBase64(proposedEncodedSignature)
	if signatureDecodeError != nil {
		return 0, "", InvalidTokenError
	}

	signer := hmac.New(crypto.SHA256.New, *privateKey)
	// sign the incoming encoded token
	signer.Write(proposedEncodedToken)
	actualSignature := signer.Sum(nil)

	// and see if the signature they have matches what we would produce
	if !hmac.Equal(proposedSignature, actualSignature) {
		return 0, "", UnauthorizedError
	}

	proposedSerializedToken, tokenDecodeError := decodeBase64(proposedEncodedToken)
	if tokenDecodeError != nil {
		return 0, "", InvalidTokenError
	}

	var successfulToken AuthToken
	if jsoniter.Unmarshal(proposedSerializedToken, &successfulToken) != nil {
		return 0, "", InvalidTokenError
	}

	// check the expires
	if successfulToken.Expires() <= time.Now().Unix() {
		return 0, "", ExpiredTokenError
	}

	internalSlug := successfulToken.InternalSlug()

	hashId, hashIdError := hashids.NewWithData(HashIDData)
	userIdArray, decodeError := hashId.DecodeInt64WithError(internalSlug)
	if hashIdError != nil || decodeError != nil || len(userIdArray) != 1 {
		return 0, "", InvalidTokenError
	}

	return userIdArray[0], internalSlug, nil
}


var passwordHasher argon2.Config = argon2.DefaultConfig()
func hashPassword(password *[]byte) ([]byte, error) {
	encodedPassword, encodingError := passwordHasher.HashEncoded(*password)
	argon2.SecureZeroMemory(*password)
	return encodedPassword, encodingError
}

func verifyPassword(encodedPassword *[]byte, trialPassword *[]byte) bool {
	match, err := argon2.VerifyEncoded(*trialPassword, *encodedPassword)
	argon2.SecureZeroMemory(*encodedPassword)
	argon2.SecureZeroMemory(*trialPassword)
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
