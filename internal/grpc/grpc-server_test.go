package grpc

import (
	"fmt"
	"net"
	"testing"

	"github.com/TakenokoTech/grpc-go/configs"
)

var (
	server = GServer{}
)

func TestSuccessServer(t *testing.T) {
	configs.GrpcPort = emptyPort()
	server.Start()
	server.Stop()
}

func TestFailedServer(t *testing.T) {
	configs.GrpcPort = 999999
	server.Start()
	server.Stop()
}

func emptyPort() int {
	for port := 3000; port < 20000; port++ {
		l, err := net.Listen("tcp", fmt.Sprintf("localhost:%d", port))
		if err == nil {
			defer l.Close()
			return port
		}
	}
	return 3000
}
