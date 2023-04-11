package awss3

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/vesicash/upload-ms/internal/config"
	"github.com/vesicash/upload-ms/utility"
)

var (
	S3Session *session.Session
)

func ConnectAws(logger *utility.Logger) *session.Session {
	awss3config := config.GetConfig().AWSS3
	AccessKeyID := awss3config.AccessKeyId
	SecretAccessKey := awss3config.SecretAccessKey
	MyRegion := awss3config.DefaultRegion
	sess, err := session.NewSession(
		&aws.Config{
			Region: aws.String(MyRegion),
			Credentials: credentials.NewStaticCredentials(
				AccessKeyID,
				SecretAccessKey,
				"",
			),
		})
	if err != nil {
		logger.Error("creating s3 session", err)
		panic(err)
	}
	logger.Info("created s3 session")
	S3Session = sess
	return sess
}
