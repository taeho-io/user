package handler

import (
	"database/sql"
	"strings"

	"github.com/taeho-io/auth"
	"github.com/taeho-io/go-taeho/id"
	"github.com/taeho-io/user"
	"github.com/taeho-io/user/server/models"
	"github.com/volatiletech/sqlboiler/boil"
	. "github.com/volatiletech/sqlboiler/queries/qm"
	"golang.org/x/net/context"
	"google.golang.org/api/oauth2/v2"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type SignInWithGoogleHandlerFunc func(ctx context.Context, request *user.SignInWithGoogleRequest) (*user.SignInWithGoogleResponse, error)

func SignInWithGoogle(oauth2Svc *oauth2.Service, id id.ID, db *sql.DB, authCli auth.AuthClient) SignInWithGoogleHandlerFunc {
	return func(ctx context.Context, req *user.SignInWithGoogleRequest) (*user.SignInWithGoogleResponse, error) {
		if err := req.Validate(); err != nil {
			return nil, status.Error(codes.InvalidArgument, err.Error())
		}

		tokenInfoCall := oauth2Svc.Tokeninfo()
		tokenInfoCall.IdToken(req.GoogleIdToken)
		tokenInfo, err := tokenInfoCall.Do()
		if err != nil {
			return nil, status.Error(codes.Unauthenticated, err.Error())
		}

		var userID int64
		u, err := models.Users(Where("email=?", tokenInfo.Email)).One(ctx, db)
		switch err {
		case nil:
			userID = u.ID
		case sql.ErrNoRows:
			newUserID, err := id.Generate()
			if err != nil {
				return nil, status.Error(codes.Internal, err.Error())
			}

			userID = newUserID

			u := &models.User{
				ID:             userID,
				Type:           user.UserType_GOOGLE.String(),
				Email:          tokenInfo.Email,
				HashedPassword: "",
				Name:           req.Name,
			}

			err = u.Insert(ctx, db, boil.Infer())
			if err != nil {
				if strings.Contains(err.Error(), idxEmailUniqueViolation) {
					return nil, ErrEmailAlreadyExists
				}
				return nil, status.Error(codes.Internal, err.Error())
			}
		default:
			return nil, status.Error(codes.Internal, err.Error())
		}

		authResp, err := authCli.Auth(ctx, &auth.AuthRequest{
			UserId: userID,
		})
		if err != nil {
			return nil, err
		}

		return &user.SignInWithGoogleResponse{
			AccessToken:  authResp.AccessToken,
			RefreshToken: authResp.RefreshToken,
			ExpiresIn:    authResp.ExpiresIn,
			UserId:       userID,
		}, nil
	}
}
