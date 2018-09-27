package main

import (
	"fmt"
	"regexp"

	"gopkg.in/guregu/null.v3"
	"github.com/gin-gonic/gin"
	"github.com/blainehansen/goqu"

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


type uploadResponse struct {
	ObjectName string
	Signature string
	Timestamp int64
}

type uploadConfirmation struct {
	Signature string `binding:"required"`
	Timestamp int64 `binding:"required"`
	Hash string `binding:"required"`
	Version string `binding:"required"`
}

var imagesProfilePreset string = environment["CDN_API_PROFILE_IMAGES_PRESET"]

var _ r = authRoute(POST, "/user/profile-image/sign", func(c *gin.Context) {
	userSlug := c.MustGet("userSlug").(string)

	// signature, timestamp := SignUploadParams(userSlug, imagesProfilePreset)
	objectName := fmt.Sprintf("%s.mp3", userSlug)
	signature, timestamp := SignUploadParams(objectName, imagesProfilePreset)

	response := uploadResponse {
		ObjectName: objectName,
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


var imagesProjectPreset string = environment["CDN_API_PROJECT_IMAGES_PRESET"]

func checkUserOwnsProject(c *gin.Context, userId int64, projectId int64) bool {
	var count int64
	found, err := Projects.Table.Where(
		Projects.Id.Eq(projectId), Projects.UserId.Eq(userId),
	).Select(
		goqu.COUNT("*"),
	).ScanVal(&count)
	if err != nil {
		c.AbortWithError(500, err)
		return false
	}
	if !found {
		c.AbortWithError(500, fmt.Errorf("count not found? %s", count))
		return false
	}

	switch count {
		case 0:
			// obscure the fact that there's even a project here
			c.AbortWithStatus(404)
			return false
		case 1:
			return true
		default:
			c.AbortWithError(500, fmt.Errorf("count didn't make any sense: %s for projectId %s and userId %s", count, projectId, userId))
			return false
	}
}


func makeProjectUploadName(projectSlug string, hash string, version string) string {
	if version == "" {
		return fmt.Sprintf("%s/%s", projectSlug, hash)
	}
	return fmt.Sprintf("v%s/%s/%s", version, projectSlug, hash)
}

var _ r = authRoute(POST, "/project/:projectSlug/uploads/sign", func(c *gin.Context) {
	userId := c.MustGet("userId").(int64)

	projectSlug := c.Param("projectSlug")
	projectId, decodeError := decodeSlug(projectSlug)
	if decodeError != nil {
		c.AbortWithStatus(400); return
	}

	hashes := struct {
		Hashes []string `binding:"required"`
	} {}
	if err := c.ShouldBindJSON(&hashes); err != nil {
		c.AbortWithError(422, err); return
	}

	if !checkUserOwnsProject(c, userId, projectId) { return }

	response := []uploadResponse{}

	for _, hash := range hashes.Hashes {
		// TODO also there should probably be a length requirement
		if hashInvalidCharactersRegex.MatchString(hash) {
			c.AbortWithError(400, fmt.Errorf("hash doesn't match the required format: %s", hash)); return
		}

		objectName := makeProjectUploadName(projectSlug, hash, "")
		signature, timestamp := SignUploadParams(objectName, imagesProjectPreset)

		newResponse := uploadResponse {
			ObjectName: objectName,
			Signature: signature,
			Timestamp: timestamp,
		}
		response = append(response, newResponse)
	}

	c.JSON(200, response)
})

var _ r = authRoute(POST, "/project/:projectSlug/uploads/confirm", func(c *gin.Context) {
	userId := c.MustGet("userId").(int64)

	projectSlug := c.Param("projectSlug")
	projectId, decodeError := decodeSlug(projectSlug)
	if decodeError != nil {
		c.AbortWithStatus(400); return
	}

	confirmations := struct {
		Confirmations []uploadConfirmation `binding:"required"`
	} {}
	if err := c.ShouldBindJSON(&confirmations); err != nil {
		c.AbortWithError(422, err); return
	}

	finalUploads := []string{}

	for _, confirmation := range confirmations.Confirmations {
		signedObjectName := makeProjectUploadName(projectSlug, confirmation.Hash, "")
		if !VerifyUploadParamsSignature(confirmation.Signature, signedObjectName, imagesProjectPreset, confirmation.Timestamp) {
			c.AbortWithError(403, fmt.Errorf("invalid signature")); return
		}
		objectName := makeProjectUploadName(projectSlug, confirmation.Hash, confirmation.Version)

		finalUploads = append(finalUploads, objectName)
	}

	updateQuery := Projects.Query.Where(
		Projects.Id.Eq(projectId), Projects.UserId.Eq(userId),
	).Update(
		Projects.UploadImages.Set(finalUploads),
	)

	if !doExec(c, updateQuery) { return }

	c.Status(204)
})
