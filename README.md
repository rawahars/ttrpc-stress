# ttrpc-stress

## Overview
The `ttrpc-stress` test suite is designed to evaluate the performance and reliability of the `ttrpc` package. It includes stress tests that validate client-server interactions under load conditions. The test suite consists of:

- **Test Module:** `stress_test.go`
- **Runner Components:** `runner/` (client, server, and supporting modules)
- **Test Setup Script:** `build_runners.go`

This test suite implements a Matrix Test Strategy to ensure compatibility between different versions of TTRPC clients and servers. The goal is to validate interoperability between multiple versions and configurations in a parallelized, automated manner.

## Installation
Before running the tests, ensure that Go is installed on your system.

```sh
# Verify Go installation
go version
```

Clone the repository and navigate to the test directory:

```sh
git clone <repository_url>
cd ttrpc-stress
```

## Running the Tests
The test suite is self-sufficient and can be ran using a single command-
```sh
cd test
go test -v .
```

## Package Structure

### stress_test.go
The primary test file `stress_test.go` contains the test cases that validate `ttrpc` under stress conditions. It ensures reliable message passing and performance efficiency.

### Runner Components
- **`client.go`** – Handles client-side test logic.
- **`server.go`** – Implements the server logic.
- **`main.go`** – Entry point for running either the client or server.

## License
This project is licensed under the MIT License.

