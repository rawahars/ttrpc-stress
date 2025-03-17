package main

import (
	"context"
	"errors"
	"fmt"
	"log"

	"github.com/rawahars/ttrpc-stress/runner/connection"

	"github.com/containerd/ttrpc"
)

func RunServer(
	ctx context.Context,
	version string,
	addr string,
	svc string,
	method string,
) error {
	log.Printf("Running server for ttrpc %s version on %s", version, addr)

	// Listen for connections on the specified address.
	listener, err := connection.ListenConnection(addr)
	if err != nil {
		return fmt.Errorf("failed listening on %s: %w", addr, err)
	}

	server, err := ttrpc.NewServer()
	var serveFunc = serveNewPayload
	// Create the server based on the requested ttrpc version.
	switch version {
	case "v1.0.2", "v1.1.0":
		serveFunc = serveOldPayload
		break
	case "v1.2.0", "v1.2.4", "main", "latest":
		serveFunc = serveNewPayload
		break
	default:
		return fmt.Errorf("invalid version of ttrpc requested for stress testing")
	}

	if err != nil {
		return fmt.Errorf("failed creating ttrpc server: %w", err)
	}

	// Register a service and method with the server.
	server.Register(svc, map[string]ttrpc.Method{
		method: serveFunc,
	})

	// When the context is done, shut down the server.
	go func(ctx context.Context) {
		<-ctx.Done()
		log.Printf("Shutting down server...")
		_ = server.Close()
	}(ctx)

	// Serve the server and handle any errors.
	if err := server.Serve(ctx, listener); err != nil && !errors.Is(err, ErrServerClosed) {
		return fmt.Errorf("failed serving server: %w", err)
	}

	return nil
}

func serveOldPayload(_ context.Context, unmarshal func(interface{}) error) (interface{}, error) {
	req := &ProtogogoPayload{}
	// Unmarshal the request payload.
	if err := unmarshal(req); err != nil {
		log.Fatalf("failed unmarshalling request: %s", err)
	}
	id := req.Value
	// Return the same payload as the response.
	return &ProtogogoPayload{Value: id}, nil
}

func serveNewPayload(_ context.Context, unmarshal func(interface{}) error) (interface{}, error) {
	req := &ProtogoPayload{}
	// Unmarshal the request payload.
	if err := unmarshal(req); err != nil {
		log.Fatalf("failed unmarshalling request: %s", err)
	}
	id := req.Value
	// Return the same payload as the response.
	return &ProtogoPayload{Value: id}, nil
}
