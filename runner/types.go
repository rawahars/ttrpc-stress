package main

import (
	"errors"

	protogo "github.com/rawahars/ttrpc-stress/runner/payload_protogo"
	protogogo "github.com/rawahars/ttrpc-stress/runner/payload_protogogo"
)

var ErrServerClosed = errors.New("ttrpc: server closed")

type ProtogoPayload = protogo.Payload
type ProtogogoPayload = protogogo.Payload
