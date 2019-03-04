package handler

import (
	"database/sql"
	"strings"

	"github.com/taeho-io/auth"
	"github.com/taeho-io/go-taeho/id"
	"github.com/taeho-io/user"
	"github.com/taeho-io/user/pkg/crypt"
	"github.com/taeho-io/user/server/models"
	"github.com/volatiletech/sqlboiler/boil"
	"golang.org/x/net/context"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

const (
	idxEmailUniqueViolation = `duplicate key value violates unique constraint "idx_email"`
)

var (
	ErrEmailAlreadyExists = status.Error(codes.AlreadyExists, "email already exists")
)

type RegisterHandlerFunc func(ctx context.Context, request *user.RegisterRequest) (*user.RegisterResponse, error)

func Register(c crypt.Crypt, db *sql.DB, id id.ID, authCli auth.AuthClient) RegisterHandlerFunc {
	return func(ctx context.Context, req *user.RegisterRequest) (*user.RegisterResponse, error) {
		if err := req.Validate(); err != nil {
			return nil, status.Error(codes.InvalidArgument, err.Error())
		}

		userID, err := id.Generate()
		if err != nil {
			return nil, status.Error(codes.Internal, err.Error())
		}

		hashedPassword, err := c.HashPassword(req.GetPassword())
		if err != nil {
			return nil, status.Error(codes.Internal, err.Error())
		}

		u := &models.User{
			ID:             userID,
			Type:           req.UserType.String(),
			Email:          req.Email,
			HashedPassword: hashedPassword,
			Name:           req.Name,
		}

		err = u.Insert(ctx, db, boil.Infer())
		if err != nil {
			if strings.Contains(err.Error(), idxEmailUniqueViolation) {
				return nil, ErrEmailAlreadyExists
			}
			return nil, status.Error(codes.Internal, err.Error())
		}

		authResp, err := authCli.Auth(ctx, &auth.AuthRequest{
			UserId: userID,
		})
		if err != nil {
			return nil, status.Error(codes.Internal, err.Error())
		}

		return &user.RegisterResponse{
			AccessToken:  authResp.AccessToken,
			RefreshToken: authResp.RefreshToken,
			ExpiresIn:    authResp.ExpiresIn,
			UserId:       userID,
		}, nil
	}
}
