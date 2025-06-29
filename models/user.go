package models

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	ID        uint      `json:"id" gorm:"primary_key"`
	Email     string    `json:"email" gorm:"unique;not null"`
	Username  string    `json:"username"`
	Password  string    `json:"-" gorm:"not null"` // password will not be returned
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

// TableName sets the table name to "users"
func (u *User) TableName() string {
	return "users"
}

// BeforeCreate hook
func (u *User) BeforeCreate(tx *gorm.DB) error {
	u.CreatedAt = time.Now()
	u.UpdatedAt = time.Now()
	return nil
}

// BeforeUpdate hook
func (u *User) BeforeUpdate(tx *gorm.DB) error {
	u.UpdatedAt = time.Now()
	return nil
}
