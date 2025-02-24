package test

import (
	"context"
	"errors"
	"net"

	protogo "github.com/rawahars/ttrpc-stress/test/payload_protogo"
	protogogo "github.com/rawahars/ttrpc-stress/test/payload_protogogo"
)

var ErrServerClosed = errors.New("ttrpc: server closed")

type TtrpcServer interface {
	Register(string, string, func(context.Context, func(interface{}) error) (interface{}, error)) error
	Serve(context.Context, net.Listener) error
	Close() error
}

type TtrpcClient interface {
	Call(context.Context, string, string, interface{}, interface{}) error
	Close() error
}

type ProtogoPayload = protogo.Payload
type ProtogogoPayload = protogogo.Payload
