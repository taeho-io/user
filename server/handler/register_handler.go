package handler

import (
	"database/sql"
	"strings"

	"github.com/dchest/validator"
	"github.com/hashicorp/go-multierror"
	"github.com/pkg/errors"
	"github.com/taeho-io/auth"
	"github.com/taeho-io/user"
	"github.com/taeho-io/user/pkg/crypt"
	"github.com/taeho-io/user/server/models"
	"github.com/volatiletech/sqlboiler/boil"
	"github.com/xissy/kubeflake"
	"golang.org/x/net/context"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type RegisterHandlerFunc func(ctx context.Context, request *user.RegisterRequest) (*user.RegisterResponse, error)

func Register(c crypt.Crypt, db *sql.DB, authCli auth.AuthClient) RegisterHandlerFunc {
	return func(ctx context.Context, req *user.RegisterRequest) (*user.RegisterResponse, error) {
		err := validate(req)
		if err != nil {
			return nil, status.Error(codes.InvalidArgument, err.Error())
		}

		userID, err := kubeflake.New()
		if err != nil {
			return nil, err
		}

		hashedPassword, err := c.HashPassword(req.GetPassword())
		if err != nil {
			return nil, err
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
			if strings.Contains(err.Error(), `duplicate key value violates unique constraint "idx_email"`) {
				return nil, status.Error(codes.AlreadyExists, "email already exists")
			}
			return nil, err
		}

		authResp, err := authCli.Auth(ctx, &auth.AuthRequest{
			UserId: userID,
		})
		if err != nil {
			return nil, err
		}

		return &user.RegisterResponse{
			AccessToken:  authResp.AccessToken,
			RefreshToken: authResp.RefreshToken,
		}, nil
	}
}

var (
	ErrNotSupportedUserType = errors.New("not supported user_type")
	ErrInvalidEmail         = errors.New("invalid email")
	ErrInvalidName          = errors.New("invalid name")
	ErrPasswordTooShort     = errors.New("password too short")
)

func validate(req *user.RegisterRequest) error {
	var errs *multierror.Error

	if req.UserType != user.UserType_EMAIL {
		errs = multierror.Append(errs, ErrNotSupportedUserType)
		return errs.ErrorOrNil()
	}

	if !validator.IsValidEmail(req.Email) {
		errs = multierror.Append(errs, ErrInvalidEmail)
	}

	if strings.Trim(req.Name, " ") == "" {
		errs = multierror.Append(errs, ErrInvalidName)
	}

	if len(req.Password) < 6 {
		errs = multierror.Append(errs, ErrPasswordTooShort)
	}

	return errs.ErrorOrNil()
}
