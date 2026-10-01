package entities

import (
	"time"

	"github.com/google/uuid"
)

type Wallet struct {
	ID               uuid.UUID
	Email            string
	PasswordHash     string
	FirstName        string
	LastName         string
	Status           string
	EmailVerifiedAt  *time.Time
	PhoneVerifiedAt  *time.Time
	FailedLoginCount int
	LockedUntil      *time.Time
	CreatedAt        time.Time
	UpdatedAt        time.Time
	DeletedAt        *time.Time
}
