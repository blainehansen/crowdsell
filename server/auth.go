package main

import (
	"fmt"
	"time"
	"strings"
	"errors"

	"encoding/base64"
	"github.com/lhecker/argon2"
	"github.com/json-iterator/go"

	"github.com/speps/go-hashids"

	"crypto"
	"crypto/rand"
	_ "crypto/sha256"
	"crypto/hmac"
)


var signingKey *[]byte = func() *[]byte {
	key := []byte(environment["SIGNING_KEY"])
	return &key
}()

var HashIDData *hashids.HashIDData = func() *hashids.HashIDData {
	internalHashIdData := hashids.HashIDData {
		Alphabet: environment["HASHID_ALPHABET"],
		MinLength: environment["HASHID_MIN_LENGTH"],
		Salt: environment["HASHID_SALT"],
	}
	return &internalHashIdData
}()


func generateRandomToken() (string, error) {
	tokenBytes := make([]byte, 32)
	_, err := rand.Read(tokenBytes)
	if err != nil {
		return nil, err
	}

	return localEncoding.EncodeToString(tokenBytes), nil
}


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

func issueAuthTokenForId(userId int64) (string, string, error) {
	hashId, hashIdError := hashids.NewWithData(HashIDData)
	userInternalSlug, encodeError := hashId.EncodeInt64([]int64{userId})
	if hashIdError != nil || encodeError != nil {
		return "", "", encodeError
	}

	token, issueError := issueAuthToken(userInternalSlug)
	return userInternalSlug, token, issueError
}

func issueAuthToken(userInternalSlug string) (string, error) {
	// get tomorrow unix time
	tomorrow := time.Now().Add(time.Duration(24) * time.Hour).Unix()


	// create an authtoken
	token := AuthToken{ userInternalSlug, tomorrow }

	// json serialize it and base64 encode it
	serializedToken, serializationError := jsoniter.Marshal(token)
	if serializationError != nil {
		return "", serializationError
	}

	encodedToken := encodeBase64(serializedToken)

	// create a signature of it
	signer := hmac.New(crypto.SHA256.New, *signingKey)
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

	signer := hmac.New(crypto.SHA256.New, *signingKey)
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
