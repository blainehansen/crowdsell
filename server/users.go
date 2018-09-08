package main

import (
	"fmt"
	"regexp"

	"gopkg.in/guregu/null.v3"
	"github.com/gin-gonic/gin"
	// "github.com/badoux/checkmail"
	// "gopkg.in/doug-martin/goqu.v4"
	// "github.com/blainehansen/goqu"
)

var urlSlugInvalidCharactersRegex = regexp.MustCompile("[^[:alnum:]-]")
var hashInvalidCharactersRegex = regexp.MustCompile("[^[:alnum:]]")

var _ r = authRoute(PUT, "/user/slug", func(c *gin.Context) {
	userId := c.MustGet("userId").(int64)

	slugPayload := struct {
		Slug string `binding:"required"`
	}{}
	if err := c.ShouldBindJSON(&slugPayload); err != nil {
		c.AbortWithError(422, err); return
	}

	if urlSlugInvalidCharactersRegex.MatchString(slugPayload.Slug) {
		c.AbortWithError(400, fmt.Errorf("slug doesn't match the required format: %s", slugPayload.Slug)); return
	}

	updateQuery := Users.Query.Where(
		Users.Id.Eq(userId),
	).Update(
		Users.UrlSlug.Set(slugPayload.Slug),
	)

	if !doExec(c, updateQuery) { return }

	c.Status(204)
})


var _ r = authRoute(GET, "/user", func(c *gin.Context) {
	userId := c.MustGet("userId").(int64)

	user := struct {
		Name null.String
		Bio null.String
		Links null.String
		Location null.String
		ProfilePhotoVersion null.String
	}{}
	found, err := Users.Query.Where(
		Users.Id.Eq(userId),
	).ScanStruct(&user)

	if err != nil {
		c.AbortWithError(500, err); return
	}
	if !found {
		c.AbortWithStatus(404); return
	}

	c.JSON(200, &user)
})


var _ r = authRoute(PATCH, "/user", func(c *gin.Context) {
	userId := c.MustGet("userId").(int64)

	var userMap map[string]interface{}
	if err := c.ShouldBindJSON(&userMap); err != nil {
		c.AbortWithError(422, err); return
	}

	patchQuery := Users.Query.Where(
		Users.Id.Eq(userId),
	).Patch(userMap)

	if !doPatch(c, patchQuery) { return }

	c.Status(204)
})



// func makeProfileObjectName(userSlug string) string {
// 	return fmt.Sprintf("%s.png", userSlug)
// }

var imagesProfilePreset string = environment["CDN_API_PROFILE_IMAGES_PRESET"]

var _ r = authRoute(POST, "/user/profile-image/sign", func(c *gin.Context) {
	userSlug := c.MustGet("userSlug").(string)

	signature, timestamp := SignUploadParams(userSlug, imagesProfilePreset)

	response := struct {
		ObjectName string
		Signature string
		Timestamp int64
	} {
		ObjectName: userSlug,
		Signature: signature,
		Timestamp: timestamp,
	}
	c.JSON(200, response)
})


var _ r = authRoute(POST, "/user/profile-image/confirm", func(c *gin.Context) {
	userId := c.MustGet("userId").(int64)
	userSlug := c.MustGet("userSlug").(string)

	confirmationPayload := struct {
		Signature string `binding:"required"`
		Timestamp int64 `binding:"required"`
		Version string `binding:"required"`
	}{}
	if err := c.ShouldBindJSON(&confirmationPayload); err != nil {
		c.AbortWithError(422, err); return
	}

	if !VerifyUploadParamsSignature(confirmationPayload.Signature, userSlug, imagesProfilePreset, confirmationPayload.Timestamp) {
		c.AbortWithError(403, fmt.Errorf("invalid signature")); return
	}

	updateQuery := Users.Query.Where(
		Users.Id.Eq(userId),
	).Update(
		Users.ProfilePhotoVersion.Set(confirmationPayload.Version),
	)
	if !doExec(c, updateQuery) { return }

	c.Status(204)
})
