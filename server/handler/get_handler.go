package handler

import (
	"github.com/taeho-io/user"
	"golang.org/x/net/context"
)

type GetHandlerFunc func(ctx context.Context, request *user.GetRequest) (*user.GetResponse, error)

func Get() GetHandlerFunc {
	return func(ctx context.Context, req *user.GetRequest) (*user.GetResponse, error) {
		return &user.GetResponse{}, nil
	}
}
