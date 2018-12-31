package handler

import (
	"database/sql"
	"testing"
	"time"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"github.com/taeho-io/user"
	"golang.org/x/net/context"
	"gopkg.in/DATA-DOG/go-sqlmock.v1"
)

func TestGet_Success(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	ctx := context.Background()
	db, sqlMock, err := sqlmock.New()
	assert.Nil(t, err)
	rows := sqlmock.NewRows([]string{"id", "type ", "email", "hashed_password", "name", "created_at", "updated_at"}).
		AddRow(userID, userType, email, hashedPassword, name, time.Now(), time.Now())
	sqlMock.ExpectQuery(`^SELECT \* FROM \"taeho\".\"user\".*`).
		WillReturnRows(rows)

	resp, err := Get(db)(ctx, &user.GetRequest{
		UserId: userID,
	})
	assert.NotNil(t, resp)
	assert.Nil(t, err)
}

func TestGet_Validate_Error(t *testing.T) {
	ctx := context.Background()

	db, _, _ := sqlmock.New()
	resp, err := Get(db)(ctx, &user.GetRequest{
		UserId: 0,
	})
	assert.Nil(t, resp)
	assert.Error(t, err)
}

func TestGet_UserNotFound_Error(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	ctx := context.Background()
	db, sqlMock, err := sqlmock.New()
	assert.Nil(t, err)
	sqlMock.ExpectQuery(`^SELECT \* FROM \"taeho\".\"user\".*`).
		WillReturnError(sql.ErrNoRows)

	resp, err := Get(db)(ctx, &user.GetRequest{
		UserId: userID,
	})
	assert.Nil(t, resp)
	assert.EqualError(t, err, ErrUserNotFound.Error())
}

func TestGet_DB_Error(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	ctx := context.Background()
	db, sqlMock, err := sqlmock.New()
	assert.Nil(t, err)
	sqlMock.ExpectQuery(`^SELECT \* FROM \"taeho\".\"user\".*`).
		WillReturnError(sql.ErrConnDone)

	resp, err := Get(db)(ctx, &user.GetRequest{
		UserId: userID,
	})
	assert.Nil(t, resp)
	assert.Error(t, err)
}
