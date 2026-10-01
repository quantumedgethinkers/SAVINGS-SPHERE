package services

import "golang.org/x/crypto/bcrypt"

type PasswordService struct{}

func NewPasswordService() *PasswordService {
	return &PasswordService{}
}

func (p *PasswordService) Hash(password string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword(
		[]byte(password),
		bcrypt.DefaultCost,
	)
	if err != nil {
		return "", err
	}

	return string(hash), nil
}

func (p *PasswordService) Compare(hash, password string) error {
	return bcrypt.CompareHashAndPassword(
		[]byte(hash),
		[]byte(password),
	)
}
