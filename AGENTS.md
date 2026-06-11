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

### Environment update script

The update script in `.cursor/environment.json` validates the build on VM
startup. It also bootstraps `go.mod` if the file is missing, so module-mode
Go commands work without manual setup.
