package user

import (
	"sync"

	"google.golang.org/grpc"
)

var (
	cm     = &sync.Mutex{}
	Client UserClient
)

func GetUserClient() UserClient {
	cm.Lock()
	defer cm.Unlock()

	if Client != nil {
		return Client
	}

	serviceURL := "user:80"

	// We don't need to error here, as this creates a pool and connection
	// will happen later
	conn, _ := grpc.Dial(
		serviceURL,
		grpc.WithInsecure(),
	)

	cli := NewUserClient(conn)
	Client = cli
	return cli
}
