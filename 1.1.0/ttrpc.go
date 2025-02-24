package v1_1_0

import (
	"context"
	"fmt"
	"net"

	"github.com/containerd/ttrpc"
)

type Client struct {
	client *ttrpc.Client
}

func NewClient(conn net.Conn) *Client {
	return &Client{
		client: ttrpc.NewClient(conn),
	}
}

func (c *Client) Call(ctx context.Context, svc string, method string, req, resp interface{}) error {
	return c.client.Call(ctx, svc, method, req, resp)
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
	unmarshal func(context.Context, func(interface{}) error) (interface{}, error),
) error {
	if unmarshal == nil {
		return fmt.Errorf("unmarshal function is required")
	}

	s.server.Register(name, map[string]ttrpc.Method{
		method: unmarshal,
	})

	return nil
}

func (s *Server) Serve(ctx context.Context, listener net.Listener) error {
	return s.server.Serve(ctx, listener)
}

func (s *Server) Close() error {
	return s.server.Close()
}
