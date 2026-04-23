package services

import (
	"errors"
	"mime/multipart"

	"example.com/src/common/config"
	"example.com/src/models"
	"example.com/src/storages"
)

func StorageUploadPublic(file multipart.File, fileHeader multipart.FileHeader) (*models.File, error) {

	payloadFile, err := storages.UploadPublic(file, &fileHeader)

	if err != nil {
		return nil, errors.New(err.Error())
	}

	if err := config.DB.Create(payloadFile).Error; err != nil {
		return nil, errors.New(err.Error())
	}

	return payloadFile, nil
}
