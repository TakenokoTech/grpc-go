package internal

import (
	"context"
	"log"
	"time"

	pb "github.com/TakenokoTech/grpc-go/api/proto"
	"google.golang.org/grpc"
)

const (
	address     = "localhost:3000"
	defaultName = "world"
)

// Client :
type Client struct {
	conn         *grpc.ClientConn
	sampleDomain pb.SampleDomainClient
}

// Connect :
func (client *Client) Connect() {
	var err error

	client.conn, err = grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}

	client.sampleDomain = pb.NewSampleDomainClient(client.conn)
}

// Disconnect :
func (client *Client) Disconnect() {
	defer client.conn.Close()
}

// GetText :
func (client *Client) GetText(text string) error {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	r, err := client.sampleDomain.GetText(ctx, &pb.TextRequest{Text: text})
	if err != nil {
		// log.Fatalf("could not greet: %v", err)
		return err
	}

	log.Printf("Reponse: %s", r.GetText())
	return nil
}
