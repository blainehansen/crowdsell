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


var _ r = authRoute(POST, "/user/profile-image/:imageHash/:imageType", func(c *gin.Context) {
	userId := c.MustGet("userId").(int64)
	userSlug := c.MustGet("userSlug")

	file, parseErr := c.FormFile("file")
	if parseErr != nil {
		c.AbortWithError(400, parseErr); return
	}

	fileInternal, openErr := file.Open()
	if openErr != nil {
		c.AbortWithError(500, openErr); return
	}

	imageType := c.Param("imageType")
	switch imageType {
		case "png", "jpeg":
			break
		default:
			c.AbortWithStatus(400); return
	}
	imageHash := c.Param("imageHash")

	objectName := fmt.Sprintf("profile-images/%s/%s.%s", userSlug, imageHash, imageType)
	mimeType := fmt.Sprintf("image/%s", imageType)
	uploadErr := UploadToSpace(fileInternal, objectName, mimeType, UploadParams{})

	if uploadErr != nil {
		c.AbortWithError(500, uploadErr); return
	}

	result, err := Users.Query.Where(
		Users.Id.Eq(userId),
	).Update(
		Users.ProfilePhotoSlug.Set(objectName),
	).Exec()
	if err != nil {
		c.AbortWithError(500, err); return
	}
	if rowsAffected, err := result.RowsAffected(); rowsAffected == 0 || err != nil {
		c.AbortWithStatus(404); return
	}

	c.JSON(200, objectName)
})

