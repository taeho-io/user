package handler

import (
	"database/sql"

	"github.com/taeho-io/user"
	"github.com/taeho-io/user/server/models"
	"github.com/volatiletech/sqlboiler/queries/qm"
	"golang.org/x/net/context"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type GetHandlerFunc func(ctx context.Context, request *user.GetRequest) (*user.GetResponse, error)

func Get(db *sql.DB) GetHandlerFunc {
	return func(ctx context.Context, req *user.GetRequest) (*user.GetResponse, error) {
		u, err := models.Users(qm.Where("id=?", req.UserId)).One(ctx, db)
		if err != nil {
			switch err {
			case sql.ErrNoRows:
				return nil, status.Error(codes.NotFound, "no user")
			}

			return nil, err
		}

		return &user.GetResponse{
			UserId:   u.ID,
			UserType: user.UserType_EMAIL,
			Email:    u.Email,
			Name:     u.Name,
		}, nil
	}
}
