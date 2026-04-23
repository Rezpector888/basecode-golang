package config

import (
	"fmt"

	"example.com/src/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDatabase() {
	host := DatabaseConfig.Host
	port := DatabaseConfig.Port
	user := DatabaseConfig.User
	password := DatabaseConfig.Password
	name := DatabaseConfig.Name

	dsn := fmt.Sprintf("host=%v user=%v password=%v dbname=%v port=%v",
		host, user, password, name, port)

	database, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		LogMessage("ERROR", "Postgresql connection error: "+err.Error())
	}
	DB = database
	DB.AutoMigrate(
		&models.User{},
		&models.File{},
	)
}
