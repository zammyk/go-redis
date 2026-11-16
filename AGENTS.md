# Repository Guidelines

## Project Structure & Module Organization

This repository is a small Redis clone using Go’s standard library. All source files belong to the root `main` package:

- `tcp_listener.go`: TCP entry point and command dispatch on port 6379.
- `resp.go`: Redis Serialization Protocol (RESP) parsing and the `Value` representation.
- `serializer.go`: RESP response serialization and writing.
- `command_handler.go`: command registration and handlers for PING, SET, GET, HSET, and HGET.
- `ds.go`: in-memory string/hash maps and their read/write mutexes.

`go.mod` declares module `github.com/zammyk/go-redis` and Go 1.22.3. There are currently no test files, asset directories, or external dependencies.

## Build, Test, and Development Commands

Run commands from the repository root:

- `go build -o /tmp/go-redis .`: compile the server without creating a repository artifact.
- `go run .`: start the server on `:6379`; ensure the port is available.
- `go test ./...`: run all package tests once tests are added.
- `go test -race ./...`: check tests for shared-state races.
- `go vet ./...`: run Go’s static checks.
- `gofmt -w *.go`: format Go source files.

For a manual smoke check, start the server, then run `redis-cli -p 6379 PING` if Redis CLI is installed. The current listener accepts one connection and exits after that client disconnects.

## Coding Style & Naming Conventions

Use standard Go formatting with tab indentation through `gofmt`. Keep filenames lowercase with underscores, exported identifiers in PascalCase, and internal functions in mixedCaps. Register command names in uppercase in `Handlers`. Validate argument counts before accessing arguments and protect shared maps with their corresponding mutexes.

## Testing Guidelines

Use Go’s standard `testing` package. Place tests beside implementation files as `*_test.go`, with functions named `TestXxx`. Prefer table-driven cases for RESP parsing, serialization, invalid argument counts, and missing keys. Reset shared maps between handler tests; avoid parallel tests that mutate global storage. No coverage threshold is configured.

## Commit & Pull Request Guidelines

Existing commits use short imperative subjects, such as `Add HSET and HGET commands`. Follow that style and keep changes focused. Pull requests should describe the behavior change, link relevant issues, and report validation commands and results. For protocol changes, include example requests and expected RESP responses.
