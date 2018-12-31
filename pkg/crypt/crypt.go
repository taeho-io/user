package crypt

import "golang.org/x/crypto/bcrypt"

type Crypt interface {
	HashPassword(password string) (string, error)
	IsValidPassword(hashedPassword, password string) bool
}

type BCrypt struct {
	Crypt

	cfg Config
}

func New(cfg Config) Crypt {
	return &BCrypt{
		cfg: cfg,
	}
}

func Mock() Crypt {
	return New(NewMockConfig())
}

func (c *BCrypt) HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(bytes), nil
}

func (c *BCrypt) IsValidPassword(hashedPassword, password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
	return err == nil
}
