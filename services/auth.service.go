package services

import (
	"errors"
	"time"

	"example.com/common/config"
	"example.com/models"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

func AuthLogin(email, password string) (*models.User, error) {

	var user models.User

	if err := config.DB.Where("email = ?", email).First(&user).Error; err != nil {
		return nil, errors.New("email or password is wrong")
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
		return nil, errors.New("email or password is wrong")
	}

	return &user, nil
}

var jwtSecret = []byte(config.JwtConfig.Secret)

func AuthCreateToken(user *models.User) (string, error) {

	claims := jwt.MapClaims{
		"user_id":   user.ID,
		"full_name": user.Fullname,
		"email":     user.Email,
		"exp":       time.Now().Add(24 * time.Hour).Unix(),
		"iat":       time.Now().Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(jwtSecret)
}
