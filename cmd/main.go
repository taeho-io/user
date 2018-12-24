package main

import (
	"fmt"

	"github.com/taeho-io/user/server"
)

func main() {
	fmt.Println("Starting User gRPC server...")
	err := server.Serve()
	fmt.Println(err)
}
