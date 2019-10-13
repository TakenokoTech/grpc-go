package grpc

import (
	"testing"
)

func TestSuccessConnect(t *testing.T) {
	client := GClient{}
	client.Connect()
	client.Disconnect()
}
