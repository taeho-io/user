package handler

import (
	"database/sql"

	"github.com/taeho-io/user"
	"github.com/taeho-io/user/pkg/crypt"
	"golang.org/x/net/context"
)

type LogInHandlerFunc func(ctx context.Context, request *user.LogInRequest) (*user.LogInResponse, error)

func LogIn(c crypt.Crypt, db *sql.DB) LogInHandlerFunc {
	return func(ctx context.Context, req *user.LogInRequest) (*user.LogInResponse, error) {
		return &user.LogInResponse{}, nil
	}
}
