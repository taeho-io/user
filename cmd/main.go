package main

import (
	"fmt"
	"os"
	"sync"

	"github.com/taeho-io/user/server"
)

func main() {
	var wg sync.WaitGroup
	wg.Add(1)

	go func() {
		defer wg.Done()

		fmt.Println("Starting User gRPC server on :80")
		err := server.Serve(":80")
		fmt.Println(err)
	}()

	go func() {
		defer wg.Done()

		fmt.Println("Starting User gRPC server on :81")
		err := server.Serve(":81")
		fmt.Println(err)
	}()

	wg.Wait()
	os.Exit(1)
}
