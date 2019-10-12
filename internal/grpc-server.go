package internal

import (
	"context"
	"flag"
	"log"
	"net"

	pb "github.com/TakenokoTech/grpc-go/api/proto"
	"google.golang.org/grpc"
)

const (
	port = ":3000"
)

// SampleService :
type SampleService struct {
}

// GetText :
func (s *SampleService) GetText(ctx context.Context, in *pb.TextRequest) (*pb.TextReply, error) {
	log.Println("GetText")
	return &pb.TextReply{Text: "again " + in.GetText()}, nil
}

// StartServer :
func StartServer() (err error) {
	log.Println("StartServer")
	flag.Parse()

	lis, err := net.Listen("tcp", port)
	if err != nil {
		return
	}

	grpcServer := grpc.NewServer()
	pb.RegisterSampleDomainServer(grpcServer, &SampleService{})

	if flag.Lookup("test.v") == nil {
		grpcServer.Serve(lis)
	}

	return
}
