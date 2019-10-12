package main

import (
	"testing"

	"github.com/TakenokoTech/grpc-go/internal"
)

func Test_main(t *testing.T) {
	go func() {
		main()
	}()

	client := internal.Client{}
	client.Connect()
	client.GetText("__(><)__")
	client.Disconnect()
}
