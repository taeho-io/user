package handler

import (
	"database/sql"
	"testing"
	"time"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"github.com/taeho-io/auth"
	"github.com/taeho-io/idl/gen/go/user"
	"github.com/taeho-io/user/pkg/crypt"
	"golang.org/x/net/context"
	"gopkg.in/DATA-DOG/go-sqlmock.v1"
)

func TestLogIn_Success(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	c := crypt.NewMockCrypt(ctrl)
	c.EXPECT().IsValidPassword(hashedPassword, password).Return(true)
	ctx := context.Background()
	authCli := mockAuthClient(ctrl, ctx, userID)

	db, sqlMock, err := sqlmock.New()
	assert.Nil(t, err)
	rows := sqlmock.NewRows([]string{"id", "type ", "email", "hashed_password", "name", "created_at", "updated_at"}).
		AddRow(userID, userType, email, hashedPassword, name, time.Now(), time.Now())
	sqlMock.ExpectQuery(`^SELECT \* FROM \"taeho\".\"user\".*`).
		WillReturnRows(rows)

	resp, err := LogIn(c, db, authCli)(ctx, &user.LogInRequest{
		UserType: userType,
		Email:    email,
		Password: password,
	})
	assert.NotNil(t, resp)
	assert.Nil(t, err)
}

func TestLogIn_Validate_Error(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	c := crypt.Mock()
	authCli := &auth.MockAuthClient{}

	db, sqlMock, err := sqlmock.New()
	assert.Nil(t, err)
	assert.NotNil(t, sqlMock)

	ctx := context.Background()
	resp, err := LogIn(c, db, authCli)(ctx, &user.LogInRequest{
		UserType: 999,
		Email:    email,
		Password: password,
	})
	assert.Nil(t, resp)
	assert.Error(t, err)
}

func TestLogIn_UserNotFound_Error(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	c := crypt.Mock()
	authCli := &auth.MockAuthClient{}

	db, sqlMock, err := sqlmock.New()
	assert.Nil(t, err)
	sqlMock.ExpectQuery(`^SELECT \* FROM \"taeho\".\"user\".*`).
		WillReturnError(sql.ErrNoRows)

	ctx := context.Background()
	resp, err := LogIn(c, db, authCli)(ctx, &user.LogInRequest{
		UserType: userType,
		Email:    email,
		Password: password,
	})
	assert.Nil(t, resp)
	assert.NotNil(t, err)
	assert.EqualError(t, err, ErrLogInFailed.Error())
}
