package config

import (
	"fmt"
	"log"

	"github.com/Reyshal/task-manager-api/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDatabase() {
	var err error
	config := ConfigInstance.Database

	// Connect to the database
	log.Println("🚀 Connecting to database...")
	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=%s TimeZone=%s",
		config.Host, config.User, config.Password, config.DBName, config.Port, config.SSLMode, config.TimeZone,
	)

	if DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{}); err != nil {
		log.Fatal("❌ Failed to connect to database: " + err.Error())
	}

	// Migrate the database
	if err = DB.AutoMigrate(&models.User{}, &models.Task{}); err != nil {
		log.Fatal("❌ Failed to migrate database: " + err.Error())
	}

	log.Println("✅ Connected to database & migrated successfully")
}

func GetDatabase() *gorm.DB {
	return DB
}
