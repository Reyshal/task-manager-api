package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	Env      string
	Database DatabaseConfig
	JWT      JWTConfig
	Server   ServerConfig
}

type DatabaseConfig struct {
	Host     string
	User     string
	Password string
	DBName   string
	Port     string
	SSLMode  string
	TimeZone string
}

type JWTConfig struct {
	SecretKey string
}

type ServerConfig struct {
	Port string
}

var ConfigInstance *Config

func InitConfig() {
	// Load .env file
	log.Println("🚀 Loading .env file...")
	if err := godotenv.Load(); err != nil {
		log.Println("❌ Failed to load .env file: " + err.Error())
	}

	ConfigInstance = &Config{
		Env: os.Getenv("ENV"),
		Database: DatabaseConfig{
			Host:     os.Getenv("DB_HOST"),
			User:     os.Getenv("DB_USER"),
			Password: os.Getenv("DB_PASSWORD"),
			DBName:   os.Getenv("DB_NAME"),
			Port:     os.Getenv("DB_PORT"),
			SSLMode:  os.Getenv("DB_SSLMODE"),
			TimeZone: os.Getenv("DB_TIMEZONE"),
		},
		JWT: JWTConfig{
			SecretKey: os.Getenv("JWT_SECRET"),
		},
		Server: ServerConfig{
			Port: os.Getenv("SERVER_PORT"),
		},
	}

	log.Println("✅ Loaded .env file successfully")
}
