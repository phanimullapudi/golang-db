package user

import (
	"time"

	"gorm.io/gorm"
)

type Service struct {
	DB *gorm.DB
}

// equals
type User struct {
	ID        uint `gorm:"primaryKey"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
	Name      string
}

func NewService(db *gorm.DB) *Service {
	return &Service{
		DB: db,
	}
}
