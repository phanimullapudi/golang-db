package user

import (
	"time"

	"gorm.io/gorm"
)

// Service - the struct for our user service
type Service struct {
	DB *gorm.DB
}

// User - defines our user structure
type User struct {
	ID        uint `gorm:"primaryKey"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
	Name      string
}

// NewService - returns a new user service
func NewService(db *gorm.DB) *Service {
	return &Service{
		DB: db,
	}
}
