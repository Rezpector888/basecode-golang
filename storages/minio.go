package storages

import (
	"context"
	"fmt"
	"mime/multipart"
	"path/filepath"
	"strings"

	"example.com/common/config"
	"example.com/models"
	"github.com/google/uuid"
	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
)

var minioClient *minio.Client

func InitMinio() error {
	var err error

	minioClient, err = minio.New(config.MinioConfig.EndPoint, &minio.Options{
		Secure: config.MinioConfig.UseSSL,
		Creds:  credentials.NewStaticV4(config.MinioConfig.AccessKey, config.MinioConfig.SecretKey, ""),
	})
	if err != nil {
		config.LogMessage("ERROR", err.Error())
		return err
	}
	return nil
}

func generateMetadata(originalname string) (id, path, name string) {
	ext := filepath.Ext(originalname)
	fileName := strings.TrimSuffix(originalname, ext)
	name = fileName + ext
	id = uuid.New().String()
	path = fmt.Sprintf("%s/%s", id, name)
	return id, path, name
}

func UploadPublic(file multipart.File, fileHeader *multipart.FileHeader) (*models.File, error) {
	if err := InitMinio(); err != nil {
		return nil, err
	}
	defer file.Close()

	id, objectPath, fileName := generateMetadata(fileHeader.Filename)
	publicPath := "public/" + objectPath

	ctx := context.Background()

	_, err := minioClient.PutObject(ctx,
		config.MinioConfig.BucketName,
		publicPath,
		file,
		fileHeader.Size,
		minio.PutObjectOptions{
			ContentType: fileHeader.Header.Get("Content-Type"),
		},
	)
	if err != nil {
		return nil, fmt.Errorf("failed to upload: %w", err)
	}

	scheme := "http"
	if config.MinioConfig.UseSSL {
		scheme = "https"
	}
	url := fmt.Sprintf("%s://%s/%s/%s", scheme, config.MinioConfig.EndPoint, config.MinioConfig.BucketName, publicPath)

	return &models.File{
		ID:       id,
		Name:     fileName,
		Path:     publicPath,
		Url:      url,
		Mimetype: fileHeader.Header.Get("Content-Type"),
	}, nil
}
