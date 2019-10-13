package main

import (
	"fmt"

	"github.com/TakenokoTech/grpc-go/internal/grpc"
)

var (
	server = grpc.GServer{}
)

func main() {
	fmt.Println("=============================")
	fmt.Println("TakenokoTech/grpc-go")
	fmt.Println("=============================")
	server.Start()
}
