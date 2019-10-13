package main

import (
	"testing"

	"github.com/TakenokoTech/grpc-go/internal/grpc"
)

var (
	server = grpc.GServer{}
)

func Test_main(t *testing.T) {
	go func() {
		server.Start()
	}()
	main()
}
