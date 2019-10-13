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
	client.SampleDomain.GetText("(((^-^)))")
	client.SampleDomain.GetText("__(>__<)__")
	client.SampleDomain.GetText("o(*^▽^*)o")
	client.SampleDomain.GetText("＼(^o^)／")
	client.Disconnect()
}
