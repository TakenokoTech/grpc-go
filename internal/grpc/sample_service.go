package grpc

import (
	"context"
	"log"

	pb "github.com/TakenokoTech/grpc-go/api/proto"
)

// SampleService :
type SampleService struct {
}

// GetText :
func (s *SampleService) GetText(ctx context.Context, in *pb.TextRequest) (*pb.TextReply, error) {
	log.Println("GetText")
	return &pb.TextReply{Text: "again " + in.GetText()}, nil
}
