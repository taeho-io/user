package handler

import (
	"database/sql"

	"github.com/taeho-io/auth"
	"github.com/taeho-io/user"
	"github.com/taeho-io/user/pkg/crypt"
	"github.com/taeho-io/user/server/models"
	"github.com/volatiletech/sqlboiler/queries/qm"
	"golang.org/x/net/context"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type LogInHandlerFunc func(ctx context.Context, request *user.LogInRequest) (*user.LogInResponse, error)

var (
	ErrLogInFailed = status.Error(codes.Unauthenticated, "logIn failed")
)

func LogIn(c crypt.Crypt, db *sql.DB, authCli auth.AuthClient) LogInHandlerFunc {
	return func(ctx context.Context, req *user.LogInRequest) (*user.LogInResponse, error) {
		if err := req.Validate(); err != nil {
			return nil, status.Error(codes.InvalidArgument, err.Error())
		}

		u, err := models.Users(qm.Where("email=?", req.Email)).One(ctx, db)
		if err != nil {
			return nil, ErrLogInFailed
		}

		if !c.IsValidPassword(u.HashedPassword, req.Password) {
			return nil, ErrLogInFailed
		}

		authResp, err := authCli.Auth(ctx, &auth.AuthRequest{
			UserId: u.ID,
		})
		if err != nil {
			return nil, err
		}

		return &user.LogInResponse{
			AccessToken:  authResp.AccessToken,
			RefreshToken: authResp.RefreshToken,
			UserId:       authResp.UserId,
			ExpiresIn:    authResp.ExpiresIn,
		}, nil
	}
}
