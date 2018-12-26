package handler

import (
	"testing"

	"github.com/hashicorp/go-multierror"
	"github.com/stretchr/testify/assert"
	"github.com/taeho-io/user"
)

func TestValidate_ErrNotSupportedUserType(t *testing.T) {
	req := &user.RegisterRequest{
		UserType: 999,
	}
	err := validate(req)
	assert.NotNil(t, err)
	assert.Equal(t, multierror.Append(ErrNotSupportedUserType), err)
}

func TestValidate_ErrInvalidEmail(t *testing.T) {
	req := &user.RegisterRequest{
		UserType: user.UserType_EMAIL,
		Email:    "taeho@taeho.",
		Name:     "Taeho",
		Password: "ad8hfa82b#s8",
	}
	err := validate(req)
	assert.NotNil(t, err)
	assert.Equal(t, multierror.Append(ErrInvalidEmail), err)
}

func TestValidate_ErrInvalidName(t *testing.T) {
	req := &user.RegisterRequest{
		UserType: user.UserType_EMAIL,
		Email:    "taeho@taeho.io",
		Name:     "",
		Password: "ad8hfa82b#s8",
	}
	err := validate(req)
	assert.NotNil(t, err)
	assert.Equal(t, multierror.Append(ErrInvalidName), err)
}

func TestValidate_ErrPasswordTooShort(t *testing.T) {
	req := &user.RegisterRequest{
		UserType: user.UserType_EMAIL,
		Email:    "taeho@taeho.io",
		Name:     "Taeho",
		Password: "1234",
	}
	err := validate(req)
	assert.NotNil(t, err)
	assert.Equal(t, multierror.Append(ErrPasswordTooShort), err)
}

func TestValidate_ErrMultiError(t *testing.T) {
	req := &user.RegisterRequest{
		UserType: user.UserType_EMAIL,
		Email:    "taeho@taeho.",
		Name:     "",
		Password: "1234",
	}
	err := validate(req)
	assert.NotNil(t, err)
	assert.Equal(t, multierror.Append(
		ErrInvalidEmail,
		ErrInvalidName,
		ErrPasswordTooShort,
	), err)
}

func TestValidate(t *testing.T) {
	req := &user.RegisterRequest{
		UserType: user.UserType_EMAIL,
		Email:    "taeho@taeho.io",
		Name:     "Taeho",
		Password: "ad8hfa82b#s8",
	}
	err := validate(req)
	assert.Nil(t, err)
}
