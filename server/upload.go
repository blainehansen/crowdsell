package main

import (
	"io"
	"fmt"
	"strconv"

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
var shouldUpload bool = func() bool {
	innerBool, parseError := strconv.ParseBool(environment["SHOULD_UPLOAD"])
	if parseError != nil {
		panic(parseError)
	}

	return innerBool
}()

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

	if shouldUpload {
		_, err := S3Client.PutObject(&object)
		return err
	}
	return nil
}
