//go:build windows
// +build windows

package test

import (
	"net"

	"github.com/Microsoft/go-winio"
)

// listenConnection listens for incoming named pipe connections at the specified address.
func listenConnection(addr string) (net.Listener, error) {
	return winio.ListenPipe(addr, &winio.PipeConfig{
		// 0 buffer sizes for pipe is important to help deadlock to occur.
		// It can still occur if there is buffering, but it takes more IO volume to hit it.
		InputBufferSize:  0,
		OutputBufferSize: 0,
	})
}

// dialConnection dials a named pipe connection to the specified address.
func dialConnection(addr string) (net.Conn, error) {
	return winio.DialPipe(addr, nil)
}
