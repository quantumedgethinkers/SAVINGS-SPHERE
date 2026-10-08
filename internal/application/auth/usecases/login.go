package usecases

import (
	"context"
	"errors"
	"strings"

	"github.com/quantumedgethinkers/SAVINGS-SPHERE/internal/application/auth/dto"
	"github.com/quantumedgethinkers/SAVINGS-SPHERE/internal/application/auth/services"
	"github.com/quantumedgethinkers/SAVINGS-SPHERE/internal/domain/repositories"
	"github.com/quantumedgethinkers/SAVINGS-SPHERE/internal/jwt"
)

var (
	ErrInvalidCredentials = errors.New("invalid email or password")
)

type LoginUseCase struct {
	userRepository  repositories.UserRepository
	passwordService *services.PasswordService
	jwtService      *jwt.JWTService
}

func NewLoginUseCase(
	userRepository repositories.UserRepository,
	passwordService *services.PasswordService,
	jwtService *jwt.JWTService,
) *LoginUseCase {
	return &LoginUseCase{
		userRepository:  userRepository,
		passwordService: passwordService,
		jwtService:      jwtService,
	}
}

func (u *LoginUseCase) Execute(
	ctx context.Context,
	req dto.LoginRequest,
) (*dto.LoginResponse, error) {

	email := strings.ToLower(
		strings.TrimSpace(req.Email),
	)

	user, err := u.userRepository.FindByEmail(
		ctx,
		email,
	)
	if err != nil {
		return nil, err
	}

	if user == nil {
		return nil, ErrInvalidCredentials
	}

	err = u.passwordService.Compare(
		user.PasswordHash,
		req.Password,
	)
	if err != nil {
		return nil, ErrInvalidCredentials
	}

	accessToken, expiresIn, err :=
		u.jwtService.GenerateAccessToken(
			user.ID.String(),
		)

	if err != nil {
		return nil, err
	}

	return &dto.LoginResponse{
		AccessToken: accessToken,
		TokenType:   "Bearer",
		ExpiresIn:   expiresIn,
	}, nil
}
