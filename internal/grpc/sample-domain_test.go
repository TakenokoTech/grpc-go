package grpc

import (
	"context"
	"testing"

	pb "github.com/TakenokoTech/grpc-go/api/proto"
)

func TestSampleDomain_GetText(t *testing.T) {
	service := &SampleDomain{}
	r, err := service.GetText(context.Background(), &pb.TextRequest{Text: "test"})

	if err != nil || r.Text == "" {
		t.Errorf("error")
	}
}
