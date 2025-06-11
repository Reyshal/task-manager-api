package models

import (
	"time"

	"gorm.io/gorm"
)

type Priority string

const (
	PriorityLow    Priority = "low"
	PriorityMedium Priority = "medium"
	PriorityHigh   Priority = "high"
)

type Task struct {
	ID          uint       `json:"id" gorm:"primary_key"`
	Title       string     `json:"title" gorm:"not null"`
	Description string     `json:"description"`
	Completed   bool       `json:"completed" gorm:"default:false"`
	Priority    Priority   `json:"priority" gorm:"default:low"`
	DueDate     *time.Time `json:"due_date,omitempty"`
	UserID      uint       `json:"user_id" gorm:"not null"`
	User        User       `json:"user,omitempty" gorm:"foreignKey:UserID"`
	CreatedAt   time.Time  `json:"created_at"`
	UpdatedAt   time.Time  `json:"updated_at"`
}

// TableName sets the table name to "tasks"
func (t *Task) TableName() string {
	return "tasks"
}

// BeforeCreate hook
func (t *Task) BeforeCreate(tx *gorm.DB) error {
	t.CreatedAt = time.Now()
	t.UpdatedAt = time.Now()
	return nil
}

// BeforeUpdate hook
func (t *Task) BeforeUpdate(tx *gorm.DB) error {
	t.UpdatedAt = time.Now()
	return nil
}
