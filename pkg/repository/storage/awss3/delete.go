package awss3

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/vesicash/upload-ms/internal/config"
)

var (
	DeleteObjectInTestMode = ""
)

func DeleteUpload(fileName string) error {
	awss3config := config.GetConfig().AWSS3
	svc := s3.New(S3Session)
	input := &s3.DeleteObjectInput{
		Bucket: aws.String(awss3config.Bucket),
		Key:    aws.String(fileName),
	}

	_, err := svc.DeleteObject(input)
	if err != nil {
		return err
		// if aerr, ok := err.(awserr.Error); ok {
		// 	switch aerr.Code() {
		// 	default:
		// 		fmt.Println(aerr.Error())
		// 	}
		// } else {
		// 	// Print the error, cast err to awserr.Error to get the Code and
		// 	// Message from an error.
		// 	fmt.Println(err.Error())
		// }
		// return
	}
	return nil
}
