package awss3

import (
	"mime/multipart"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
	"github.com/vesicash/upload-ms/internal/config"
)

func UploadToS3(fileName string, file multipart.File) (string, error) {
	var (
		uploader    = s3manager.NewUploader(S3Session)
		awss3config = config.GetConfig().AWSS3
		myBucket    = awss3config.Bucket
	)

	up, err := uploader.Upload(&s3manager.UploadInput{
		Bucket: aws.String(myBucket),
		ACL:    aws.String("public-read"),
		Key:    aws.String(fileName),
		Body:   file,
	})

	if err != nil {
		return "", err
	}

	DeleteObjectInTestMode = fileName

	return up.Location, nil
}
