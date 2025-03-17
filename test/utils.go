package main

import (
	"fmt"
	"runtime"

	"github.com/google/uuid"
)

// getFilename generates a filename for the runner binary based on the given version.
// Appends .exe if running on Windows.
func getFilename(version string) string {
	filename := fmt.Sprintf("runner_%s", version)
	if runtime.GOOS == "windows" {
		filename = fmt.Sprintf("%s.exe", filename)
	}
	return filename
}

// getTestBinaryPath returns the full path to the test binary based on the version.
func getTestBinaryPath(version string) string {
	return fmt.Sprintf("%s%s", binDir, getFilename(version))
}

// getSocketAddress generates a unique socket address for communication between client and server.
// Uses Windows named pipes or Unix domain sockets based on the operating system.
func getSocketAddress() string {
	id := uuid.New().String() // Generate a unique identifier.
	if runtime.GOOS == "windows" {
		return fmt.Sprintf("\\\\.\\pipe\\ttrpc_test_%s", id) // Windows named pipe address.
	}
	return fmt.Sprintf("/tmp/ttrpc_test_%s", id) // Unix domain socket address.
}
