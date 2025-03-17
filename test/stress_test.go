package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"testing"
)

// Define binaries to be tested as clients.
var testClients = []string{"latest", "main"}

// TestMain sets up and tears down the test environment.
func TestMain(m *testing.M) {
	// Perform setup i.e. build the required executables for the test.
	if err := buildExecutables(); err != nil {
		log.Printf("Failed during setup: %v", err)
		os.Exit(1)
	}

	// Run the tests
	code := m.Run()

	// Perform cleanup after all tests have run.
	if err := cleanup(); err != nil {
		log.Printf("Failed during teardown: %v", err)
		os.Exit(1)
	}

	// Exit with the test result code
	os.Exit(code)
}

// TestTTRPCMatrix runs tests for various combinations of client and server versions.
func TestTTRPCMatrix(t *testing.T) {
	for _, client := range testClients {
		for _, server := range ttrpcVersions {

			t.Run(fmt.Sprintf("Client_%s_Server_%s", client, server), func(t *testing.T) {
				t.Parallel() // Run this test in parallel
				if err := runTest(client, server); err != nil {
					t.Errorf("Test failed for Client: %s, Server: %s - Error: %v", client, server, err)
				}
			})

			t.Run(fmt.Sprintf("Client_%s_Server_%s", server, client), func(t *testing.T) {
				t.Parallel() // Run the opposite test in parallel
				if err := runTest(server, client); err != nil {
					t.Errorf("Test failed for Client: %s, Server: %s - Error: %v", server, client, err)
				}
			})
		}
	}
}

// runTest executes a client-server test using the given versions.
func runTest(clientVersion, serverVersion string) error {
	fmt.Printf("Running test: Client [%s] -> Server [%s]\n", clientVersion, serverVersion)

	addr := getSocketAddress() // Get the socket address for communication.

	testClientBinaryPath := getTestBinaryPath(clientVersion) // Path to client binary.
	testServerBinaryPath := getTestBinaryPath(serverVersion) // Path to server binary.

	// Start the server process.
	serverCmd := exec.Command(testServerBinaryPath, "--mode", "server", "--version", serverVersion, "--addr", addr)
	serverCmd.Stdout = os.Stdout
	serverCmd.Stderr = os.Stderr

	if err := serverCmd.Start(); err != nil {
		return fmt.Errorf("server failed to start: %w", err)
	}
	defer serverCmd.Process.Kill() // Ensure cleanup of server process.

	// Start the client process and connect to the server.
	clientCmd := exec.Command(testClientBinaryPath, "--mode", "client", "--version", clientVersion, "--addr", addr)
	clientCmd.Stdout = os.Stdout
	clientCmd.Stderr = os.Stderr

	if err := clientCmd.Run(); err != nil {
		return fmt.Errorf("client failed: %w", err)
	}

	return nil
}
