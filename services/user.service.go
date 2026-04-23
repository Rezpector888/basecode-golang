package services

import (
	"errors"
	"strings"

	"example.com/common/config"
	"example.com/controllers/user/dto"
	"example.com/models"
	"golang.org/x/crypto/bcrypt"
)

func UserFindById(userId string) (*models.User, error) {
	var user models.User

	if err := config.DB.Where("id = ?", userId).First(&user).Error; err != nil {
		return nil, errors.New("user not found")
	}

	return &user, nil

}

func UserFindAll(search string, limit, offset int) ([]models.User, int, error) {
	var users []models.User

	var total int64

	query := config.DB.Model(&models.User{})

	if search != "" {
		query = query.Where("LOWER(full_name) LIKE ?", "%"+strings.ToLower(search)+"%")
	}

	if err := query.Count(&total).Error; err != nil {
		return nil, 0, err
	}

	if err := query.Limit(limit).Offset(offset).Find(&users).Error; err != nil {
		return nil, 0, err
	}

	return users, int(total), nil
}

func UserCreate(payload dto.CreateUserInput) (*models.User, error) {

	var existing models.User

	if err := config.DB.Where("LOWER(email) = ?", strings.ToLower(payload.Email)).First(&existing).Error; err == nil {
		return nil, errors.New("email is already registered")
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(payload.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, errors.New("failed to hash password")
	}

	user := models.User{
		Fullname: payload.Fullname,
		Email:    payload.Email,
		Password: string(hashedPassword),
	}

	if err := config.DB.Create(&user).Error; err != nil {
		return nil, err
	}

	user.Password = ""
	return &user, nil

}

func UserUpdate(userId string, payload dto.UpdateUserInput) (*models.User, error) {
	var user models.User
	if err := config.DB.First(&user, "id = ?", userId).Error; err != nil {
		return nil, errors.New("user not found")
	}
	if payload.Email != "" && !strings.EqualFold(payload.Email, user.Email) {
		var existing models.User
		if err := config.DB.Where("LOWER(email) = ? AND id != ?", strings.ToLower(payload.Email), userId).
			First(&existing).Error; err == nil {
			return nil, errors.New("email already in use")
		}
		user.Email = payload.Email
	}

	if payload.Fullname != "" {
		user.Fullname = payload.Fullname
	}

	if err := config.DB.Save(&user).Error; err != nil {
		return nil, err
	}

	return &user, nil
}

func UserDelete(userId string) (*models.User, error) {
	var user models.User

	if err := config.DB.First(&user, "id = ?", userId).Error; err != nil {
		return nil, errors.New("user not found")
	}

	if err := config.DB.Delete(&user).Error; err != nil {
		return nil, err
	}

	return &user, nil
}
