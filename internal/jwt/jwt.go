package jwt

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type JWTService struct {
	Secret string
}

func New(secret string) *JWTService {
	return &JWTService{
		Secret: secret,
	}
}

func (j *JWTService) GenerateAccessToken(userID string) (string, error) {
	claims := jwt.MapClaims{
		"sub":  userID,
		"exp":  time.Now().Add(15 * time.Minute).Unix(),
		"iat":  time.Now().Unix(),
		"type": "access",
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	return token.SignedString([]byte(j.Secret))
}
