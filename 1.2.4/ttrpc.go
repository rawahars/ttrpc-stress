package v1_2_4

import (
	"context"
	"log"
	"net"

	"github.com/ttrpc-stress/payload/protogo"

	"github.com/containerd/ttrpc"
)

type payload = protogo.Payload

type Client struct {
	client *ttrpc.Client
}

func NewClient(conn net.Conn) *Client {
	return &Client{
		client: ttrpc.NewClient(conn),
	}
}

func (c *Client) Call(ctx context.Context, svc string, method string, reqId uint32) (uint32, error) {
	req := &payload{Value: reqId}
	resp := &payload{}
	err := c.client.Call(ctx, svc, method, req, resp)
	if err != nil {
		return 0, err
	}
	return resp.Value, nil
}

func (c *Client) Close() error {
	return c.client.Close()
}

type Server struct {
	server *ttrpc.Server
}

func NewServer() (*Server, error) {
	server, err := ttrpc.NewServer()
	if err != nil {
		return nil, err
	}
	return &Server{
		server: server,
	}, nil
}

func (s *Server) Register(
	name string,
	method string,
) error {
	s.server.Register(name, map[string]ttrpc.Method{
		method: func(ctx context.Context, unmarshal func(interface{}) error) (interface{}, error) {
			req := &payload{}
			// Unmarshal the request payload.
			if err := unmarshal(req); err != nil {
				log.Fatalf("failed unmarshalling request: %s", err)
			}
			// Return the same payload as the response.
			return &payload{Value: req.Value}, nil
		},
	})
	return nil
}

func (s *Server) Serve(ctx context.Context, listener net.Listener) error {
	return s.server.Serve(ctx, listener)
}

func (s *Server) Close() error {
	return s.server.Close()
}
