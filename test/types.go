package test

import (
	"context"
	"errors"
	"net"
)

var ErrServerClosed = errors.New("ttrpc: server closed")

type TtrpcServer interface {
	Register(string, string) error
	Serve(context.Context, net.Listener) error
	Close() error
}

type TtrpcClient interface {
	Call(context.Context, string, string, uint32) (uint32, error)
	Close() error
}
