//go:build !windows
// +build !windows

package connection

import "net"

// ListenConnection listens for incoming Unix domain socket connections at the specified address.
func ListenConnection(addr string) (net.Listener, error) {
	return net.Listen("unix", addr)
}

// DialConnection dials a Unix domain socket connection to the specified address.
func DialConnection(addr string) (net.Conn, error) {
	return net.Dial("unix", addr)
}
