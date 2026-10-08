package jwt

import (
	"errors"
	"strconv"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

var (
	ErrInvalidToken = errors.New("invalid token")
)

type JWTService struct {
	Secret string
}

func New(secret string) *JWTService {
	return &JWTService{
		Secret: secret,
	}
}

func (j *JWTService) GenerateAccessToken(
	userID string,
) (string, int64, error) {

	now := time.Now()
	expiresAt := now.Add(15 * time.Minute)

	claims := jwt.MapClaims{
		"sub":  userID,
		"iat":  now.Unix(),
		"exp":  expiresAt.Unix(),
		"type": "access",
	}

	token := jwt.NewWithClaims(
		jwt.SigningMethodHS256,
		claims,
	)

	tokenString, err := token.SignedString(
		[]byte(j.Secret),
	)
	if err != nil {
		return "", 0, err
	}

	expiresIn := int64(time.Until(expiresAt).Seconds())

	return tokenString, expiresIn, nil
}

func (j *JWTService) ValidateToken(
	tokenString string,
) (jwt.MapClaims, error) {

	token, err := jwt.Parse(
		tokenString,
		func(token *jwt.Token) (interface{}, error) {

			if token.Method != jwt.SigningMethodHS256 {
				return nil, ErrInvalidToken
			}

			return []byte(j.Secret), nil
		},
	)

	if err != nil {
		return nil, ErrInvalidToken
	}

	if !token.Valid {
		return nil, ErrInvalidToken
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return nil, ErrInvalidToken
	}

	tokenType, ok := claims["type"].(string)
	if !ok || tokenType != "access" {
		return nil, ErrInvalidToken
	}

	return claims, nil
}

func GetUserID(claims jwt.MapClaims) (string, error) {
	sub, ok := claims["sub"]
	if !ok {
		return "", ErrInvalidToken
	}

	switch value := sub.(type) {
	case string:
		return value, nil

	case float64:
		return strconv.FormatFloat(
			value,
			'f',
			-1,
			64,
		), nil

	default:
		return "", ErrInvalidToken
	}
}
