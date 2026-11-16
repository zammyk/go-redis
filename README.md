# go-redis

A small Redis clone built from scratch in Go, using only the standard library. It parses Redis Serialization Protocol (RESP) requests over TCP and stores strings and hashes in memory.

## Requirements

- Go 1.22.3 or later
- Optional: `redis-cli` for interactive use

## Run locally

```sh
git clone https://github.com/zammyk/go-redis.git
cd go-redis
go run .
```

The server listens on port `6379` on all network interfaces. Ensure the port is available before starting it.

To build and run a binary instead:

```sh
go build -o /tmp/go-redis .
/tmp/go-redis
```

## Usage

In another terminal, open one interactive client session:

```sh
redis-cli -h 127.0.0.1 -p 6379
```

Run these commands within that same session:

```text
127.0.0.1:6379> PING
PONG
127.0.0.1:6379> SET greeting hello
OK
127.0.0.1:6379> GET greeting
"hello"
127.0.0.1:6379> HSET user:1 name Ada
OK
127.0.0.1:6379> HGET user:1 name
"Ada"
```

The server accepts a single client connection and exits when that client disconnects. Restart it before opening another session.

## Supported commands

| Command | Behavior |
| --- | --- |
| `PING [message]` | Returns `PONG`, or echoes the first argument. |
| `SET key value` | Stores a string and returns `OK`. |
| `GET key` | Returns the stored string, or null for a missing key. |
| `HSET hash field value` | Stores one hash field and returns `OK`. |
| `HGET hash field` | Returns the field value, or null if absent. |

Command names are case-insensitive. `COMMAND` currently maps to the PING handler; command introspection is not implemented.

## Project structure

- `tcp_listener.go`: server entry point, connection handling, and command dispatch.
- `resp.go`: RESP request parsing and value representation.
- `serializer.go`: RESP response encoding and writing.
- `command_handler.go`: command registry and handlers.
- `ds.go`: shared in-memory maps and mutexes.

## Development

```sh
gofmt -w *.go
go vet ./...
go test ./...
go test -race ./...
```

No automated tests are currently included. See [AGENTS.md](AGENTS.md) for contributor guidelines.

## Current limitations

This is a learning project with a partial Redis implementation:

- Data is lost when the process exits; persistence and expiration are not implemented.
- Multiple client connections are not supported.
- Unsupported commands can panic the server, and malformed requests are not robustly handled.
- `HSET` accepts only one field/value pair and returns `OK` rather than Redis's integer result.
- Authentication and TLS are not implemented. Run it only in a trusted development environment.
