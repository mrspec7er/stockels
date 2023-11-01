package utils

import (
	"bytes"
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/feature/s3/manager"
	"github.com/aws/aws-sdk-go-v2/service/s3"
)

func FileUploader(fileStreamData *bytes.Buffer, fileName string) (*manager.UploadOutput, error) {
	cfg, err := config.LoadDefaultConfig(context.TODO())
	if err != nil {
		return &manager.UploadOutput{}, err
	}
	client := s3.NewFromConfig(cfg)
	uploader := manager.NewUploader(client)
	result, err := uploader.Upload(context.TODO(), &s3.PutObjectInput{
		Bucket: aws.String("stockels"),
		Key:    aws.String(fileName),
		Body:   fileStreamData,
		ACL: 	"public-read",
	})

	return result, err
}