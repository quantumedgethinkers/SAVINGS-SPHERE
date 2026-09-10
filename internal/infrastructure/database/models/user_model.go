package models

import (
	"time"

	"github.com/google/uuid"
)

type UserModel struct {
	ID               uuid.UUID `gorm:"type:uuid;default:gen_random_uuid();primaryKey"`
	Email            string    `gorm:"uniqueIndex;not null"`
	PasswordHash     string    `gorm:"not null"`
	FirstName        string    `gorm:"not null"`
	LastName         string    `gorm:"not null"`
	Status           string    `gorm:"default:active"`
	EmailVerifiedAt  *time.Time
	PhoneVerifiedAt  *time.Time
	FailedLoginCount int
	LockedUntil      *time.Time
	CreatedAt        time.Time
	UpdatedAt        time.Time
	DeletedAt        *time.Time `gorm:"index"`
}

func (UserModel) TableName() string {
	return "users"
}
