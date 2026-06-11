# AGENTS.md

## Cursor Cloud specific instructions

This is a single, tiny Go HTTP server (`main.go`) that listens on port `8080` and returns the "Hello from Docker!" banner. It uses only the Go standard library — there are no external dependencies and no test files.

### Key gotcha: no `go.mod`

The repo predates Go modules and has **no `go.mod`**, so default module-mode Go commands fail with `cannot find main module`. You MUST prefix Go commands with `GO111MODULE=off` to run them in legacy GOPATH mode:

- Build: `GO111MODULE=off make build` (binary → `bin/single-dev-env`) or `GO111MODULE=off go build`
- Run (dev): `GO111MODULE=off make run` (i.e. `go run main.go`)
- Vet/lint: `GO111MODULE=off go vet ./...` (no dedicated linter is configured)
- Test: `GO111MODULE=off go test ./...` (currently reports `[no test files]`)

Do not add a `go.mod` just to make commands work; use the `GO111MODULE=off` prefix instead.

### Running / verifying the server

Start the server (long-running) in a tmux session, then verify with curl:

```
curl http://localhost:8080
```

A `200` response with the Docker whale banner means the environment is working. The handler also prints the request's raw query string to stdout, so a request like `http://localhost:8080?hello=world` logs `hello=world` in the server terminal.
