package handler

import (
	"database/sql"

	"github.com/taeho-io/auth"
	"github.com/taeho-io/user"
	"github.com/taeho-io/user/server/models"
	. "github.com/volatiletech/sqlboiler/queries/qm"
	"golang.org/x/net/context"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

var (
	ErrUserNotFound     = status.Error(codes.NotFound, "user not found")
	ErrPermissionDenied = status.Error(codes.PermissionDenied, "forbidden")
)

type GetHandlerFunc func(ctx context.Context, req *user.GetRequest) (*user.GetResponse, error)

func Get(db *sql.DB) GetHandlerFunc {
	return func(ctx context.Context, req *user.GetRequest) (*user.GetResponse, error) {
		if err := req.Validate(); err != nil {
			return nil, status.Error(codes.InvalidArgument, err.Error())
		}

		if err := auth.VerifyUser(ctx, req.UserId, true); err != nil {
			return nil, ErrPermissionDenied
		}

		u, err := models.Users(Where("id=?", req.UserId)).One(ctx, db)
		if err != nil {
			switch err {
			case sql.ErrNoRows:
				return nil, ErrUserNotFound
			}

			return nil, status.Error(codes.Internal, err.Error())
		}

		return &user.GetResponse{
			UserId:   u.ID,
			UserType: user.UserType_EMAIL,
			Email:    u.Email,
			Name:     u.Name,
		}, nil
	}
}
