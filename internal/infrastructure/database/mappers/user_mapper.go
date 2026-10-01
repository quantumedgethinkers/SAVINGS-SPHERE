package mappers

import (
	"github.com/quantumedgethinkers/SAVINGS-SPHERE/internal/domain/entities"
	"github.com/quantumedgethinkers/SAVINGS-SPHERE/internal/infrastructure/database/models"
)

func ToUserModel(user *entities.User) *models.UserModel {
	return &models.UserModel{
		ID:               user.ID,
		Email:            user.Email,
		PasswordHash:     user.PasswordHash,
		FirstName:        user.FirstName,
		LastName:         user.LastName,
		Status:           user.Status,
		EmailVerifiedAt:  user.EmailVerifiedAt,
		PhoneVerifiedAt:  user.PhoneVerifiedAt,
		FailedLoginCount: user.FailedLoginCount,
		LockedUntil:      user.LockedUntil,
		CreatedAt:        user.CreatedAt,
		UpdatedAt:        user.UpdatedAt,
		DeletedAt:        user.DeletedAt,
	}
}

func ToUserEntity(model *models.UserModel) *entities.User {
	return &entities.User{
		ID:               model.ID,
		Email:            model.Email,
		PasswordHash:     model.PasswordHash,
		FirstName:        model.FirstName,
		LastName:         model.LastName,
		Status:           model.Status,
		EmailVerifiedAt:  model.EmailVerifiedAt,
		PhoneVerifiedAt:  model.PhoneVerifiedAt,
		FailedLoginCount: model.FailedLoginCount,
		LockedUntil:      model.LockedUntil,
		CreatedAt:        model.CreatedAt,
		UpdatedAt:        model.UpdatedAt,
		DeletedAt:        model.DeletedAt,
	}
}
