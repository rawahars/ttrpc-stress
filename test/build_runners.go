package main

import (
	"fmt"
	"os"
	"os/exec"
)

// Define versions to test.
var (
	ttrpcVersions = []string{"v1.0.2", "v1.1.0", "v1.2.0", "v1.2.4", "latest", "main"} // List of ttrpc versions to test
	binDir        = "./bin/"                                                           // Directory to store built executables
)

// buildExecutables builds the runner for each specified ttrpc version.
func buildExecutables() error {
	for _, version := range ttrpcVersions {
		fmt.Printf("\nBuilding runner for ttrpc version: %s\n", version)

		if err := updateTTRPCVersion(version); err != nil {
			return err
		}

		if err := buildRunner(version); err != nil {
			return err
		}

		fmt.Printf("Build successful for version %s\n", version)
	}
	return nil
}

// cleanup removes the built executables from the bin directory.
func cleanup() error {
	fmt.Println("Cleaning up built executables...")

	for _, version := range ttrpcVersions {
		filePath := getTestBinaryPath(version)
		if err := os.Remove(filePath); err != nil && !os.IsNotExist(err) {
			return fmt.Errorf("failed to remove %s: %v", filePath, err)
		}
		fmt.Printf("Removed %s\n", filePath)
	}
	fmt.Println("Cleanup complete.")
	return nil
}

// updateTTRPCVersion updates the ttrpc version in the runner module.
func updateTTRPCVersion(version string) error {
	ttrpcVersion := fmt.Sprintf("github.com/containerd/ttrpc@%s", version)
	cmd := exec.Command("go", "get", ttrpcVersion)
	cmd.Dir = "../runner"
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	if err := cmd.Run(); err != nil {
		return fmt.Errorf("failed to fetch %s: %w", ttrpcVersion, err)
	}

	cmd = exec.Command("go", "mod", "tidy")
	cmd.Dir = "../runner"
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	if err := cmd.Run(); err != nil {
		return fmt.Errorf("failed to update %s: %w", ttrpcVersion, err)
	}

	return nil
}

// buildRunner compiles the runner binary for the specified version.
func buildRunner(version string) error {
	// Determine if we need to append .exe based on GOOS
	outputBinary := fmt.Sprintf("../test/bin/%s", getFilename(version))

	cmd := exec.Command("go", "build", "-o", outputBinary, ".")
	cmd.Dir = "../runner"
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	if err := cmd.Run(); err != nil {
		return fmt.Errorf("failed to build %s: %w", outputBinary, err)
	}

	return nil
}
