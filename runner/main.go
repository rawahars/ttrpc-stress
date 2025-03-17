package main

import (
	"context"
	"flag"
	"log"
	"time"
)

// Define command-line flags
var (
	mode       string
	version    string
	addr       string
	svc        string
	method     string
	timeout    time.Duration
	iterations int
	workers    int
)

func init() {
	flag.StringVar(&mode, "mode", "", "Mode to run: 'server' or 'client'")
	flag.StringVar(&version, "version", "latest", "TTRPC version to use")
	flag.StringVar(&addr, "addr", "\\\\.\\pipe\\ttrpc-stress", "Server address")
	flag.StringVar(&svc, "svc", "ttrpc.stress.test.v1", "Service name")
	flag.StringVar(&method, "method", "TestMethod", "Method name")
	flag.DurationVar(&timeout, "timeout", 60*time.Second, "Client timeout duration")
	flag.IntVar(&iterations, "iterations", 100000, "Number of iterations for client requests")
	flag.IntVar(&workers, "workers", 10, "Number of concurrent client workers")
}

func main() {
	flag.Parse()

	if mode == "" {
		log.Fatal("Error: --mode flag must be specified as either 'server' or 'client'")
	}

	ctx := context.Background()

	switch mode {
	case "server":
		err := RunServer(ctx, version, addr, svc, method)
		if err != nil {
			log.Fatalf("Server failed: %v", err)
		}

	case "client":
		err := RunClient(ctx, version, timeout, addr, svc, method, iterations, workers)
		if err != nil {
			log.Fatalf("Client failed: %v", err)
		}

	default:
		log.Fatalf("Invalid mode: %s. Use 'server' or 'client'", mode)
	}
}
