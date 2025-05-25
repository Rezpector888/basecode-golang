package seeders

import (
	"example.com/src/common/config"
	"example.com/src/models"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

func SeedUser() {
	password := "Qwerty123*"
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		config.LogMessage("ERROR", "Failed hash password: "+err.Error())
	}
	user := models.User{
		ID:       uuid.New().String(),
		Fullname: "Sample Fullname",
		Email:    "sample@mail.com",
		Password: string(hashedPassword),
	}

	if err := config.DB.Create(&user).Error; err != nil {
		config.LogMessage("ERROR", "Failed seeding user: "+err.Error())
	}
	config.LogMessage("SUCCESS", "Seeding user success")
}
