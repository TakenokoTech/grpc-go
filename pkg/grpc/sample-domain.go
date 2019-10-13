package grpc

import (
	"context"
	"log"
	"time"

	pb "github.com/TakenokoTech/grpc-go/api/proto"
)

// SampleDomain :
type SampleDomain struct {
	client pb.SampleDomainClient
}

// GetText :
func (s *SampleDomain) GetText(text string) (err error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	r, err := s.client.GetText(ctx, &pb.TextRequest{Text: text})
	if err != nil {
		return
	}

	log.Printf("Reponse: %s", r.GetText())
	return
}
