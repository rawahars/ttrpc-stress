package test

import (
	"context"
	"fmt"
	"log"
	"net"
	"time"

	"github.com/rawahars/ttrpc-stress/latest"
	"github.com/rawahars/ttrpc-stress/v1_0_2"
	"github.com/rawahars/ttrpc-stress/v1_1_0"
	"github.com/rawahars/ttrpc-stress/v1_2_0"
	"github.com/rawahars/ttrpc-stress/v1_2_4"

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
			conn, err = dialConnection(addr)
			if err == nil {
				log.Printf("Successfully connected client to the sever")
				break dialLoop
			}
			time.Sleep(500 * time.Millisecond)
		}
	}

	var client TtrpcClient
	// Determine the client and the send function based on the requested ttrpc version.
	switch version {
	case "1.0.2":
		client = v1_0_2.NewClient(conn)
		break
	case "1.1.0":
		client = v1_1_0.NewClient(conn)
		break
	case "1.2.0":
		client = v1_2_0.NewClient(conn)
		break
	case "1.2.4":
		client = v1_2_4.NewClient(conn)
		break
	case "latest":
		client = latest.NewClient(conn)
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
					if err = sendPayload(ctx, client, svc, method, uint32(i)); err != nil {
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
	return nil
}

// sendPayload sends a request to the server and verifies the response.
func sendPayload(
	ctx context.Context,
	client TtrpcClient,
	svc string,
	method string,
	id uint32,
) error {
	// Call the server method and handle any errors.
	ret, err := client.Call(ctx, svc, method, id)
	if err != nil {
		return err
	}

	// Verify the response matches the request.
	if ret != id {
		return fmt.Errorf("expected return value %d but got %d", id, ret)
	}
	return nil
}
