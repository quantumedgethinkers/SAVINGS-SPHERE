package usecases

import (
	"context"
	"errors"
	"strings"

	"github.com/google/uuid"

	"github.com/quantumedgethinkers/SAVINGS-SPHERE/internal/application/auth/dto"
	"github.com/quantumedgethinkers/SAVINGS-SPHERE/internal/application/auth/services"
	"github.com/quantumedgethinkers/SAVINGS-SPHERE/internal/domain/entities"
	"github.com/quantumedgethinkers/SAVINGS-SPHERE/internal/domain/repositories"
)

var ErrEmailAlreadyExists = errors.New("email already exists")

type RegisterUseCase struct {
	userRepository  repositories.UserRepository
	passwordService *services.PasswordService
}

func NewRegisterUseCase(
	userRepository repositories.UserRepository,
	passwordService *services.PasswordService,
) *RegisterUseCase {
	return &RegisterUseCase{
		userRepository:  userRepository,
		passwordService: passwordService,
	}
}

func (u *RegisterUseCase) Execute(
	ctx context.Context,
	req dto.RegisterRequest,
) (*dto.RegisterResponse, error) {

	email := strings.ToLower(strings.TrimSpace(req.Email))

	exists, err := u.userRepository.ExistsByEmail(ctx, email)
	if err != nil {
		return nil, err
	}

	if exists {
		return nil, ErrEmailAlreadyExists
	}

	passwordHash, err := u.passwordService.Hash(req.Password)
	if err != nil {
		return nil, err
	}

	user := &entities.User{
		ID:           uuid.New(),
		Email:        email,
		PasswordHash: passwordHash,
		FirstName:    strings.TrimSpace(req.FirstName),
		LastName:     strings.TrimSpace(req.LastName),
		Status:       "active",
	}

	err = u.userRepository.Create(ctx, user)
	if err != nil {
		return nil, err
	}

	return &dto.RegisterResponse{
		ID:        user.ID.String(),
		Email:     user.Email,
		FirstName: user.FirstName,
		LastName:  user.LastName,
		Status:    user.Status,
	}, nil
}
