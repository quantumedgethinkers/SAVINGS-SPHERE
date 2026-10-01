package repositories

import (
	"context"
	"errors"

	"github.com/google/uuid"
	"gorm.io/gorm"

	"github.com/quantumedgethinkers/SAVINGS-SPHERE/internal/domain/entities"
	"github.com/quantumedgethinkers/SAVINGS-SPHERE/internal/infrastructure/database/mappers"
	"github.com/quantumedgethinkers/SAVINGS-SPHERE/internal/infrastructure/database/models"
)

type UserRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{
		db: db,
	}
}

func (r *UserRepository) Create(
	ctx context.Context,
	user *entities.User,
) error {
	model := mappers.ToUserModel(user)

	return r.db.WithContext(ctx).Create(model).Error
}

func (r *UserRepository) FindByID(
	ctx context.Context,
	id uuid.UUID,
) (*entities.User, error) {
	var model models.UserModel

	err := r.db.WithContext(ctx).
		Where("id = ?", id).
		First(&model).Error

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}

		return nil, err
	}

	return mappers.ToUserEntity(&model), nil
}

func (r *UserRepository) FindByEmail(
	ctx context.Context,
	email string,
) (*entities.User, error) {
	var model models.UserModel

	err := r.db.WithContext(ctx).
		Where("email = ?", email).
		First(&model).Error

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}

		return nil, err
	}

	return mappers.ToUserEntity(&model), nil
}

func (r *UserRepository) ExistsByEmail(
	ctx context.Context,
	email string,
) (bool, error) {
	var count int64

	err := r.db.WithContext(ctx).
		Model(&models.UserModel{}).
		Where("email = ?", email).
		Count(&count).Error

	if err != nil {
		return false, err
	}

	return count > 0, nil
}
