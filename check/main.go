package main

import (
	"context"
	"fmt"
	"github.com/rawahars/ttrpc-stress/test"
	"time"
)

func main() {
	fmt.Println("Starting...")

	done := make(chan struct{})

	go func() {
		err := test.RunServer(
			context.Background(),
			"latest",
			"\\\\.\\pipe\\test",
			"SERVICE",
			"TEST")
		if err != nil {
			fmt.Printf("Error server: %v\n", err)
		}
	}()

	go func() {
		err := test.RunClient(
			context.Background(),
			"1.0.2",
			time.Second*60,
			"\\\\.\\pipe\\test",
			"SERVICE",
			"TEST",
			1000000,
			10)
		if err != nil {
			fmt.Printf("Error client: %v\n", err)
		}
		close(done)
	}()

	<-done
	fmt.Println("Test finished...")
}
