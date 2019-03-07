package server

import (
	"database/sql"
	"fmt"
	"net"
	"net/http"
	"time"

	"github.com/grpc-ecosystem/go-grpc-middleware"
	"github.com/grpc-ecosystem/go-grpc-middleware/logging/logrus"
	"github.com/grpc-ecosystem/go-grpc-middleware/recovery"
	"github.com/grpc-ecosystem/go-grpc-middleware/tags"
	_ "github.com/lib/pq"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
	"github.com/taeho-io/auth"
	"github.com/taeho-io/go-taeho/id"
	"github.com/taeho-io/idl/gen/go/user"
	"github.com/taeho-io/user/pkg/crypt"
	"github.com/taeho-io/user/server/handler"
	"golang.org/x/net/context"
	"google.golang.org/api/oauth2/v2"
	"google.golang.org/grpc"
	"google.golang.org/grpc/health"
	"google.golang.org/grpc/health/grpc_health_v1"
	"google.golang.org/grpc/reflection"
)

type UserServer struct {
	user.UserServer

	cfg       Config
	bcrypt    crypt.Crypt
	db        *sql.DB
	id        id.ID
	authCli   auth.AuthClient
	oauth2Svc *oauth2.Service
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
	for {
		err = db.Ping()
		if err != nil {
			logrus.Error(errors.Wrap(err, "db ping failed"))
			time.Sleep(time.Second * 5)
			continue
		}
		break
	}

	authCli := auth.GetAuthClient()

	oauth2Svc, err := oauth2.New(&http.Client{})
	if err != nil {
		return nil, err
	}

	return &UserServer{
		cfg:       cfg,
		bcrypt:    bcrypt,
		db:        db,
		id:        id.New(),
		authCli:   authCli,
		oauth2Svc: oauth2Svc,
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

func (s *UserServer) ID() id.ID {
	return s.id
}

func (s *UserServer) AuthClient() auth.AuthClient {
	return s.authCli
}

func (s *UserServer) OAuth2Service() *oauth2.Service {
	return s.oauth2Svc
}

func (s *UserServer) RegisterServer(srv *grpc.Server) {
	user.RegisterUserServer(srv, s)
}

func (s *UserServer) Register(ctx context.Context, req *user.RegisterRequest) (*user.RegisterResponse, error) {
	return handler.Register(s.Crypt(), s.DB(), s.ID(), s.AuthClient())(ctx, req)
}

func (s *UserServer) LogIn(ctx context.Context, req *user.LogInRequest) (*user.LogInResponse, error) {
	return handler.LogIn(s.Crypt(), s.DB(), s.AuthClient())(ctx, req)
}

func (s *UserServer) SignInWithGoogle(ctx context.Context, req *user.SignInWithGoogleRequest) (*user.SignInWithGoogleResponse, error) {
	return handler.SignInWithGoogle(s.OAuth2Service(), s.ID(), s.DB(), s.AuthClient())(ctx, req)
}

func (s *UserServer) Get(ctx context.Context, req *user.GetRequest) (*user.GetResponse, error) {
	return handler.Get(s.DB())(ctx, req)
}

func Serve(addr string, cfg Config) error {
	lis, err := net.Listen("tcp", addr)
	if err != nil {
		return err
	}

	logrusEntry := logrus.NewEntry(logrus.StandardLogger())

	grpcServer := grpc.NewServer(
		grpc_middleware.WithUnaryServerChain(
			grpc_ctxtags.UnaryServerInterceptor(
				grpc_ctxtags.WithFieldExtractor(grpc_ctxtags.CodeGenRequestFieldExtractor)),
			auth.TokenUnaryServerInterceptor(),
			grpc_logrus.UnaryServerInterceptor(logrusEntry),
			grpc_recovery.UnaryServerInterceptor(),
		),
	)

	healthServer := health.NewServer()
	grpc_health_v1.RegisterHealthServer(grpcServer, healthServer)

	userServer, err := New(cfg)
	if err != nil {
		return err
	}
	user.RegisterUserServer(grpcServer, userServer)

	reflection.Register(grpcServer)
	return grpcServer.Serve(lis)
}
