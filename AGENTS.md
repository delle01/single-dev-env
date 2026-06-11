# AGENTS.md

## Cursor Cloud specific instructions

This is a minimal Go HTTP server (`main.go`) that listens on port `8080` and
returns the "Hello from Docker!" ASCII banner. It uses only the Go standard
library (no external module dependencies).

### Services

- **HTTP server** (the only service). Run in development with `make run`
  (equivalent to `go run main.go`). It serves on `http://localhost:8080`. The
  `/` handler prints the request's raw query string to stdout, so a request
  like `curl 'http://localhost:8080/?hello=world'` logs `hello=world` in the
  server terminal.

### Build / lint / test

- Build: `make build` (outputs `bin/single-dev-env`).
- Lint/vet: `go vet ./...`.
- Test: `go test ./...` (currently there are no test files).

### Non-obvious notes

- Go 1.22 runs in module mode by default, but this repo originally had no
  `go.mod`. A `go.mod` (module `single-dev-env`) is required so that
  `make build`, `go vet ./...`, and `go test ./...` work. If `go.mod` is ever
  missing, recreate it with `go mod init single-dev-env`. (The startup update
  script already does this automatically when the file is absent.)
