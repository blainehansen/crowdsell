package main

import (
	"fmt"
	"time"
	"errors"
	"strconv"
	"strings"

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
	hashidMinLength, parseErr := strconv.ParseInt(environment["HASHID_MIN_LENGTH"], 10, 32)
	if parseErr != nil {
		panic(fmt.Sprintf("invalid int: %s", environment["HASHID_MIN_LENGTH"]))
	}
	HASHID_MIN_LENGTH := int(hashidMinLength)

	internalHashIdData := hashids.HashIDData {
		Alphabet: environment["HASHID_ALPHABET"],
		MinLength: HASHID_MIN_LENGTH,
		Salt: environment["HASHID_SALT"],
	}
	return &internalHashIdData
}()

func decodeSlug(slug string) (int64, error) {
	hashId, hashIdError := hashids.NewWithData(HashIDData)
	if hashIdError != nil {
		return 0, hashIdError
	}

	idArray, decodeError := hashId.DecodeInt64WithError(slug)
	if decodeError != nil {
		return 0, decodeError
	}
	if len(idArray) != 1 {
		return 0, fmt.Errorf(`slug ("%s") decoded into something other than one value: %#v`, slug, idArray)
	}

	return idArray[0], nil
}


func generateRandomToken() ([]byte, error) {
	tokenBytes := make([]byte, 32)
	_, err := rand.Read(tokenBytes)
	if err != nil {
		return nil, err
	}

	return encodeBase64(tokenBytes), nil
}


type AuthToken struct {
	I string
	E int64
	// R map[int]int
}
func (tok *AuthToken) Slug() string {
	return tok.I
}
func (tok *AuthToken) Expires() int64 {
	return tok.E
}
// func (tok *AuthToken) Role() string {
// 	return tok.R
// }

// func issueAuthTokenForId(userId int64) (string, string, error) {
// 	hashId, hashIdError := hashids.NewWithData(HashIDData)
// 	userSlug, encodeError := hashId.EncodeInt64([]int64{userId})
// 	if hashIdError != nil || encodeError != nil {
// 		return "", "", encodeError
// 	}

// 	token, issueError := issueAuthToken(userSlug)
// 	return userSlug, token, issueError
// }

func issueAuthToken(userSlug string) (string, error) {
	// get tomorrow unix time
	tomorrow := time.Now().Add(time.Duration(24) * time.Hour).Unix()

	// create an authtoken
	token := AuthToken{ userSlug, tomorrow }

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

	userSlug := successfulToken.Slug()
	userId, decodeError := decodeSlug(userSlug)
	if decodeError != nil {
		return 0, "", InvalidTokenError
	}

	return userId, userSlug, nil
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
