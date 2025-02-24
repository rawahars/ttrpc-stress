//go:build !windows
// +build !windows

package test

import "net"

// listenConnection listens for incoming Unix domain socket connections at the specified address.
func listenConnection(addr string) (net.Listener, error) {
	return net.Listen("unix", addr)
}

// dialConnection dials a Unix domain socket connection to the specified address.
func dialConnection(addr string) (net.Conn, error) {
	return net.Dial("unix", addr)
}
