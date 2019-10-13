package main

import (
	"fmt"

	"github.com/TakenokoTech/grpc-go/pkg/grpc"
)

var (
	client = grpc.GClient{}
)

func main() {
	fmt.Println("=============================")
	fmt.Println("TakenokoTech/grpc-go-client")
	fmt.Println("=============================")
	client.Connect()
	client.GetText("__(><)__")
	client.Disconnect()
}
