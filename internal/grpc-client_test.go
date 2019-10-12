package internal

import (
	"context"
	"errors"
	"testing"

	pb "github.com/TakenokoTech/grpc-go/api/proto"
	"google.golang.org/grpc"
)

type mockSampleDomainClient struct {
	cc *grpc.ClientConn
}

func (c *mockSampleDomainClient) GetText(ctx context.Context, in *pb.TextRequest, opts ...grpc.CallOption) (*pb.TextReply, error) {
	if in.Text == "" {
		return nil, errors.New("")
	}
	return nil, nil
}

func TestSuccessStartClient(t *testing.T) {
	client := Client{}
	client.sampleDomain = &mockSampleDomainClient{cc: nil}
	client.GetText("_(--)_ ")
}

func TestErrorStartClient(t *testing.T) {
	client := Client{}
	client.sampleDomain = &mockSampleDomainClient{cc: nil}
	client.GetText("")
}

func TestConnect(t *testing.T) {
	client := Client{}
	client.Connect()
	client.Disconnect()
}
