package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"time"

	"github.com/rawahars/ttrpc-stress/runner/connection"

	"github.com/containerd/ttrpc"
	"golang.org/x/sync/errgroup"
)

// RunClient runs the requested ttrpc client for the specified number of iterations and workers.
func RunClient(
	parentCtx context.Context,
	version string,
	timeout time.Duration,
	addr string,
	svc string,
	method string,
	iterations int,
	workers int,
) error {
	log.Printf("Running client for ttrpc version %s on %s", version, addr)

	// Create a context which gets cancelled if we encounter an error.
	ctx, cancel := context.WithTimeout(parentCtx, timeout)
	defer cancel() // Ensure cleanup

	var err error
	var conn net.Conn
dialLoop: // Labeled loop
	for {
		select {
		case <-ctx.Done():
			return fmt.Errorf("failed dialing connection to %s: %w", addr, ctx.Err())
		default:
			conn, err = connection.DialConnection(addr)
			if err == nil {
				log.Printf("Successfully connected client to the sever")
				break dialLoop
			}
			time.Sleep(500 * time.Millisecond)
		}
	}

	client := ttrpc.NewClient(conn)
	var sendFunc = sendNewPayload
	// Determine the send function based on the requested ttrpc version.
	switch version {
	case "v1.0.2", "v1.1.0":
		sendFunc = sendOldPayload
		break
	case "v1.2.0", "v1.2.4", "main", "latest":
		sendFunc = sendNewPayload
		break
	default:
		return fmt.Errorf("invalid version of ttrpc requested for stress testing")
	}

	defer client.Close()

	ch := make(chan int)
	var eg errgroup.Group

	// Start worker goroutines to send requests.
	for i := 0; i < workers; i++ {
		eg.Go(func() error {
			for {
				select {
				case <-ctx.Done():
					return ctx.Err()

				case i, ok := <-ch:
					if !ok {
						return nil
					}
					// Send the request and handle any errors.
					if err = sendFunc(ctx, client, svc, method, uint32(i)); err != nil {
						cancel() // stop all workers.
						return err
					}
				}
			}
		})
	}

	// Send iterations to the channel.
	for i := 0; i < iterations; i++ {
		select {
		case <-ctx.Done():
			client.Close()
			break
		case ch <- i:
		}
	}
	close(ch)

	// Wait for all goroutines to finish.
	if err = eg.Wait(); err != nil {
		return fmt.Errorf("client failure: %w", err)
	}

	log.Println("Finished all the iterations from the client")
	return nil
}

// sendOldPayload sends a request to the server and verifies the response.
func sendOldPayload(
	ctx context.Context,
	client *ttrpc.Client,
	svc string,
	method string,
	id uint32,
) error {
	req := &ProtogogoPayload{Value: id}
	resp := &ProtogogoPayload{}

	// Call the server method and handle any errors.
	if err := client.Call(ctx, svc, method, req, resp); err != nil {
		return err
	}

	ret := resp.Value
	// Verify the response matches the request.
	if ret != id {
		return fmt.Errorf("expected return value %d but got %d", id, ret)
	}
	return nil
}

// sendNewPayload sends a request to the server and verifies the response.
func sendNewPayload(
	ctx context.Context,
	client *ttrpc.Client,
	svc string,
	method string,
	id uint32,
) error {
	req := &ProtogoPayload{Value: id}
	resp := &ProtogoPayload{}

	// Call the server method and handle any errors.
	if err := client.Call(ctx, svc, method, req, resp); err != nil {
		return err
	}

	ret := resp.Value
	// Verify the response matches the request.
	if ret != id {
		return fmt.Errorf("expected return value %d but got %d", id, ret)
	}
	return nil
}
