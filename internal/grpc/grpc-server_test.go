package grpc

import (
	"fmt"
	"net"
	"os"
	"testing"
	"time"

	"github.com/TakenokoTech/grpc-go/configs"
)

var (
	server = GServer{}
)

func TestMain(m *testing.M) {
	println("before all...")
	tempPort := configs.GrpcPort
	code := m.Run()
	println("after all...")
	configs.GrpcPort = tempPort
	os.Exit(code)
}

func TestSuccessServer(t *testing.T) {
	configs.GrpcPort = emptyPort()
	go func() {
		time.Sleep(time.Millisecond * 500)
		server.Stop()
	}()
	server.Start()
}

func TestFailedServer(t *testing.T) {
	configs.GrpcPort = 999999
	server.Start()
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
