package user

import (
	"sync"

	"github.com/taeho-io/go-taeho/interceptor"
	"github.com/taeho-io/idl/gen/go/user"
	"google.golang.org/grpc"
)

const (
	serviceURL = "user:80"
)

var (
	cm     = &sync.Mutex{}
	Client user.UserClient
)

func GetUserClient() user.UserClient {
	cm.Lock()
	defer cm.Unlock()

	if Client != nil {
		return Client
	}

	// We don't need to error here, as this creates a pool and connection
	// will happen later
	conn, _ := grpc.Dial(
		serviceURL,
		grpc.WithInsecure(),
		grpc.WithUnaryInterceptor(
			interceptor.ContextUnaryClientInterceptor(),
		),
	)

	cli := user.NewUserClient(conn)
	Client = cli
	return cli
}
