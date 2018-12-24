package server

import (
	"database/sql"
	"fmt"
	"net"

	_ "github.com/lib/pq"
	"github.com/taeho-io/auth"
	"github.com/taeho-io/user"
	"github.com/taeho-io/user/pkg/crypt"
	"github.com/taeho-io/user/server/handler"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

type UserServer struct {
	user.UserServer

	cfg     Config
	bcrypt  crypt.Crypt
	db      *sql.DB
	authCli auth.AuthClient
}

func New(cfg Config) (*UserServer, error) {
	bcrypt := crypt.New(crypt.NewConfig())

	dsn := fmt.Sprintf(
		"host=%s dbname=%s user=%s password=%s sslmode=disable",
		cfg.Settings().PostgresHost,
		cfg.Settings().PostgresDBName,
		cfg.Settings().PostgresUser,
		cfg.Settings().PostgresPassword,
	)
	db, err := sql.Open("postgres", dsn)
	if err != nil {
		return nil, err
	}
	err = db.Ping()
	if err != nil {
		return nil, err
	}

	authCli := auth.GetAuthClient()

	return &UserServer{
		cfg:     cfg,
		bcrypt:  bcrypt,
		db:      db,
		authCli: authCli,
	}, nil
}

func Mock() *UserServer {
	s, _ := New(MockConfig())
	return s
}

func (s *UserServer) Config() Config {
	return s.cfg
}

func (s *UserServer) Crypt() crypt.Crypt {
	return s.bcrypt
}

func (s *UserServer) DB() *sql.DB {
	return s.db
}

func (s *UserServer) AuthClient() auth.AuthClient {
	return s.authCli
}

func (s *UserServer) RegisterServer(srv *grpc.Server) {
	user.RegisterUserServer(srv, s)
}

func (s *UserServer) Register(ctx context.Context, req *user.RegisterRequest) (*user.RegisterResponse, error) {
	return handler.Register(s.Crypt(), s.DB(), s.AuthClient())(ctx, req)
}

func (s *UserServer) LogIn(ctx context.Context, req *user.LogInRequest) (*user.LogInResponse, error) {
	return handler.LogIn(s.Crypt(), s.DB())(ctx, req)
}

func Serve() error {
	lis, err := net.Listen("tcp", "0.0.0.0:80")
	if err != nil {
		return err
	}

	grpcServer := grpc.NewServer()
	userServer, err := New(NewConfig(NewSettings()))
	if err != nil {
		return err
	}
	user.RegisterUserServer(grpcServer, userServer)
	reflection.Register(grpcServer)
	return grpcServer.Serve(lis)
}
