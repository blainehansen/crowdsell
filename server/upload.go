package main

import (
	"io"

	"github.com/gin-gonic/gin"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
)

var S3Client *s3.S3 = s3.New(session.New(&aws.Config{
	Credentials: credentials.NewStaticCredentials(
		// accessKeyID
		"WVTBP6TRFMDIQSRFHNQS",
		// secretAccessKey
		"ulHKld8JmkSqIoebaKVXyqZQEVc9+bODk4zblJAuZT8",
		"",
	),
	Endpoint: aws.String("https://nyc3.digitaloceanspaces.com"),
	// dumb aws hack
	Region: aws.String("us-east-1"),
}))


const bucketName string = "blaine-final-spaces-test"
func UploadToSpace(fileObject io.ReadSeeker) error {
	objectName := "user/9/profile.jpg"

	object := s3.PutObjectInput{
		Body: fileObject,
		Bucket: aws.String(bucketName),
		Key: aws.String(objectName),
		ACL: aws.String("public-read"),
		CacheControl: aws.String("public, max-age=31556926"),
		ContentEncoding: aws.String("gzip"),
		ContentDisposition: aws.String("inline"),
		ContentType: aws.String("image/jpeg"),
	}
	_, err := S3Client.PutObject(&object)
	return err
}


var _ r = route(POST, "/upload", func(c *gin.Context) {
	file, parseErr := c.FormFile("file")
	if parseErr != nil {
		c.AbortWithStatus(400)
		return
	}

	fileInternal, openErr := file.Open()
	if openErr != nil {
		c.AbortWithError(500, openErr)
		return
	}
	uploadErr := UploadToSpace(fileInternal)

	if uploadErr != nil {
		c.AbortWithError(500, uploadErr)
	} else {
		c.Status(204)
	}
})
