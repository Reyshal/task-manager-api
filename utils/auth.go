package utils

import (
	"time"

	"github.com/Reyshal/task-manager-api/config"
	"github.com/Reyshal/task-manager-api/models"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(bytes), err
}

func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

func GenerateToken(user models.User) (string, error) {
	// Create claims
	claims := jwt.MapClaims{
		"user_id":  user.ID,
		"username": user.Username,
		"exp":      jwt.NewNumericDate(time.Now().Add(time.Hour * 24)),
	}

	// Create token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Generate encoded token and return it
	t, err := token.SignedString([]byte(config.ConfigInstance.JWT.SecretKey))

	return t, err
}
