package grpc

import (
	"fmt"
	"log"

	pb "github.com/TakenokoTech/grpc-go/api/proto"
	"github.com/TakenokoTech/grpc-go/configs"
	"google.golang.org/grpc"
)

// GClient :
type GClient struct {
	conn         *grpc.ClientConn
	SampleDomain SampleDomain
}

// Connect :
func (client *GClient) Connect() (err error) {
	log.Printf("Connect: %d", configs.GrpcPort)
	client.conn, _ = grpc.Dial(fmt.Sprintf("localhost:%d", configs.GrpcPort), grpc.WithInsecure())
	client.SampleDomain.client = pb.NewSampleDomainClient(client.conn)
	return nil
}

// Disconnect :
func (client *GClient) Disconnect() {
	defer client.conn.Close()
}
