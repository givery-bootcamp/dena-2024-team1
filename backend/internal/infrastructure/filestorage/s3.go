package filestorage

import (
	"errors"
	"log"
	"mime/multipart"
	"myapp/internal/config"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
)

type FileStorage interface {
	UploadFile(file *multipart.File, fileName string) error
}

func setUpFileStorage() (*s3.S3, error) {
	cred := credentials.NewStaticCredentials(config.AwsAccessKeyID, config.AwsSecretAccessKey, "")

	newSession, err := session.NewSession(&aws.Config{
		Credentials: cred,
		Region:      aws.String(config.AwsDefaultRegion),
	})

	if err != nil {
		return nil, err
	}

	svc := s3.New(newSession)

	return svc, nil
}

type S3FileStorage struct{}

func (S3FileStorage S3FileStorage) UploadFile(file *multipart.File, fileName string) error {
	svc, err := setUpFileStorage()

	if err != nil {
		return err
	}

	key := "images/" + fileName

	// 拡張子をfileNameの最後の.以降の文字列として取得
	var ext string
	for i := len(fileName) - 1; i >= 0; i-- {
		if fileName[i] == '.' {
			ext = fileName[i+1:]
			break
		}
	}

	// 拡張子が取得できなかった場合はエラー
	if ext == "" {
		return errors.New("failed to get file extension")
	}

	params := &s3.PutObjectInput{
		Bucket:      aws.String(config.S3Bucket),
		Body:        *file,
		Key:         aws.String(key),
		ContentType: aws.String("image/" + ext),
	}

	_, err = svc.PutObject(params)
	if err != nil {
		log.Fatal(err)
	}

	return nil
}

func SetUpS3() (fileStorage FileStorage) {
	return S3FileStorage{}
}
