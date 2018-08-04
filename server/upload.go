package main

import (
	"io"
	"fmt"

	"github.com/gin-gonic/gin"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
)


var S3Client *s3.S3 = s3.New(session.New(&aws.Config{
	Credentials: credentials.NewStaticCredentials(
		// accessKeyID
		environment["SPACES_ACCESS_KEY"],
		// secretAccessKey
		environment["SPACES_SECRET_KEY"],
		// some unknown unimportant thing
		"",
	),
	Endpoint: aws.String(environment["SPACES_ENDPOINT"]),
	// dumb aws hack
	Region: aws.String("us-east-1"),
}))


type UploadParams struct {
	Private bool
	NoCache bool
	DispositionFilename string
}

var bucketName string = environment["SPACES_BUCKET_NAME"]
func UploadToSpace(fileObject io.ReadSeeker, objectKey string, contentType string, params UploadParams) error {
	var aclString string
	if params.Private {
		aclString = "private"
	} else {
		aclString = "public-read"
	}

	var cacheString string
	if params.NoCache {
		cacheString = "private, must-revalidate"
	} else {
		cacheString = "public, max-age=31556926"
	}

	var dispositionString string
	if params.DispositionFilename != "" {
		dispositionString = fmt.Sprintf(`attachment; filename="%s"`, params.DispositionFilename)
	} else {
		dispositionString = "inline"
	}

	object := s3.PutObjectInput{
		Body: fileObject,
		Bucket: aws.String(bucketName),
		Key: aws.String(objectKey),
		ACL: aws.String(aclString),
		CacheControl: aws.String(cacheString),
		// ContentEncoding: aws.String("gzip"),
		ContentDisposition: aws.String(dispositionString),
		ContentType: aws.String(contentType),
	}
	_, err := S3Client.PutObject(&object)
	return err
}

var _ r = authRoute(POST, "/profile-image/:imageHash/:imageType", func(c *gin.Context) {
	userId := c.MustGet("userId").(int64)
	userSlug := c.MustGet("userSlug")

	file, parseErr := c.FormFile("file")
	if parseErr != nil {
		c.AbortWithStatus(400); return
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

	result, err := UsersTable.Where(
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
