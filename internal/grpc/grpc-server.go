package grpc

import (
	"flag"
	"fmt"
	"log"
	"net"

	pb "github.com/TakenokoTech/grpc-go/api/proto"
	"github.com/TakenokoTech/grpc-go/configs"
	"google.golang.org/grpc"
)

//GServer :
type GServer struct {
	serve *grpc.Server
}

// Start :
func (server *GServer) Start() (err error) {
	log.Println("Start")
	flag.Parse()

	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", configs.GrpcPort))
	if err != nil {
		return
	}

	server.serve = grpc.NewServer()
	pb.RegisterSampleDomainServer(server.serve, &SampleService{})

	if flag.Lookup("test.v") == nil {
		server.serve.Serve(lis)
	}

	return
}

// Stop :
func (server *GServer) Stop() {
	log.Println("Stop")
	server.serve.Stop()
}
