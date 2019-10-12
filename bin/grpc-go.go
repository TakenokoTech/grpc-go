package main

import (
	"fmt"

	"github.com/TakenokoTech/grpc-go/internal"
)

func main() {
	fmt.Println("=============================")
	fmt.Println("TakenokoTech/grpc-go")
	fmt.Println("=============================")
	fmt.Println("")
	internal.StartServer()
}
