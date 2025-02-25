package test

import (
	"context"
	"errors"
	"fmt"
	"log"

	"github.com/rawahars/ttrpc-stress/latest"
	"github.com/rawahars/ttrpc-stress/v1_0_2"
	"github.com/rawahars/ttrpc-stress/v1_1_0"
	"github.com/rawahars/ttrpc-stress/v1_2_0"
	"github.com/rawahars/ttrpc-stress/v1_2_4"
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
	listener, err := listenConnection(addr)
	if err != nil {
		return fmt.Errorf("failed listening on %s: %w", addr, err)
	}

	var server TtrpcServer
	// Create the server based on the requested ttrpc version.
	switch version {
	case "1.0.2":
		server, err = v1_0_2.NewServer()
		break
	case "1.1.0":
		server, err = v1_1_0.NewServer()
		break
	case "1.2.0":
		server, err = v1_2_0.NewServer()
		break
	case "1.2.4":
		server, err = v1_2_4.NewServer()
		break
	case "latest":
		server, err = latest.NewServer()
		break
	default:
		return fmt.Errorf("invalid version of ttrpc requested for stress testing")
	}

	if err != nil {
		return fmt.Errorf("failed creating ttrpc server: %w", err)
	}

	// Register a service and method with the server.
	if err = server.Register(
		svc,
		method,
	); err != nil {
		return fmt.Errorf("failed registering service: %w", err)
	}

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
