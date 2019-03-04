package handler

import (
	"database/sql"
	"errors"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/icrowley/fake"
	"github.com/stretchr/testify/assert"
	"github.com/taeho-io/auth"
	"github.com/taeho-io/go-taeho/id"
	"github.com/taeho-io/user"
	"github.com/taeho-io/user/pkg/crypt"
	"golang.org/x/net/context"
	"gopkg.in/DATA-DOG/go-sqlmock.v1"
)

var (
	userType       = user.UserType_EMAIL
	email          = fake.EmailAddress()
	password       = fake.SimplePassword()
	name           = fake.FirstName()
	userID         = id.New().Must()
	hashedPassword = "hashed_password"
)

func mockID(ctrl *gomock.Controller, userID int64) id.ID {
	tid := id.NewMockID(ctrl)
	tid.EXPECT().Generate().Return(userID, nil)
	return tid
}

func mockAuthClient(ctrl *gomock.Controller, ctx context.Context, userID int64) auth.AuthClient {
	authCli := auth.NewMockAuthClient(ctrl)
	authCli.
		EXPECT().
		Auth(ctx, &auth.AuthRequest{
			UserId: userID,
		}).
		Return(&auth.AuthResponse{
			AccessToken:  "access_token",
			RefreshToken: "refresh_token",
			ExpiresIn:    3600,
			UserId:       userID,
		}, nil)

	return authCli
}

func TestRegister_Success(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	c := crypt.NewMockCrypt(ctrl)
	c.EXPECT().HashPassword(password).Return(hashedPassword, nil)
	tid := mockID(ctrl, userID)
	ctx := context.Background()
	authCli := mockAuthClient(ctrl, ctx, userID)

	db, sqlMock, err := sqlmock.New()
	assert.Nil(t, err)
	sqlMock.ExpectExec(`^INSERT INTO \"taeho\".\"user\".*`).
		WithArgs(sqlmock.AnyArg(),
			userType.String(),
			email,
			hashedPassword,
			name,
			sqlmock.AnyArg(),
			sqlmock.AnyArg()).
		WillReturnResult(sqlmock.NewResult(userID, 1))

	resp, err := Register(c, db, tid, authCli)(ctx, &user.RegisterRequest{
		UserType: userType,
		Email:    email,
		Password: password,
		Name:     name,
	})
	assert.NotNil(t, resp)
	assert.Nil(t, err)
}

func TestRegister_Validate_Error(t *testing.T) {
	c := crypt.Mock()
	db, _, _ := sqlmock.New()
	tid := id.New()
	authCli := &auth.MockAuthClient{}

	ctx := context.Background()
	resp, err := Register(c, db, tid, authCli)(ctx, &user.RegisterRequest{
		UserType: 999,
		Email:    email,
		Password: password,
		Name:     name,
	})
	assert.Nil(t, resp)
	assert.Error(t, err)
}

func TestRegister_EmailAlreadyExists_Error(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	c := crypt.NewMockCrypt(ctrl)
	c.EXPECT().HashPassword(password).Return(hashedPassword, nil)
	tid := mockID(ctrl, userID)
	ctx := context.Background()
	authCli := &auth.MockAuthClient{}

	db, sqlMock, err := sqlmock.New()
	assert.Nil(t, err)
	sqlMock.ExpectExec(`INSERT INTO "taeho"."user"`).
		WithArgs(sqlmock.AnyArg(),
			userType.String(),
			email,
			hashedPassword,
			name,
			sqlmock.AnyArg(),
			sqlmock.AnyArg()).
		WillReturnError(errors.New(`duplicate key value violates unique constraint "idx_email"`))

	resp, err := Register(c, db, tid, authCli)(ctx, &user.RegisterRequest{
		UserType: userType,
		Email:    email,
		Password: password,
		Name:     name,
	})
	assert.Nil(t, resp)
	assert.Error(t, err)
	assert.EqualError(t, err, ErrEmailAlreadyExists.Error())
}

func TestRegister_DB_Error(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	c := crypt.NewMockCrypt(ctrl)
	c.EXPECT().HashPassword(password).Return(hashedPassword, nil)
	tid := mockID(ctrl, userID)
	ctx := context.Background()
	authCli := &auth.MockAuthClient{}

	db, sqlMock, err := sqlmock.New()
	assert.Nil(t, err)
	sqlMock.ExpectExec(`INSERT INTO "taeho"."user"`).
		WillReturnError(sql.ErrConnDone)

	resp, err := Register(c, db, tid, authCli)(ctx, &user.RegisterRequest{
		UserType: userType,
		Email:    email,
		Password: password,
		Name:     name,
	})
	assert.Nil(t, resp)
	assert.Error(t, err)
}
