# Signals

A simple Go package for handling system signals and user input in a clean, channel-based way.

## Overview

The `signals` package provides easy-to-use functions that return channels which close when specific signals are received. This allows for clean signal handling in Go applications using `select` statements.

## Installation

```bash
go get github.com/nyxstack/signals
```

## Functions

### Signal Handlers

### `Interrupt() <-chan struct{}`
Returns a channel that closes when `SIGINT` (Ctrl+C) is received.

### `Terminate() <-chan struct{}`
Returns a channel that closes when `SIGTERM` is received (commonly used by Kubernetes and process managers).

### `Hangup() <-chan struct{}`
Returns a channel that closes when `SIGHUP` is received (terminal hangup, reload, or session end).

### `Sigusr1() <-chan struct{}` (Unix only)
Returns a channel that closes when `SIGUSR1` is received (user-defined signal 1).

### `Sigusr2() <-chan struct{}` (Unix only)
Returns a channel that closes when `SIGUSR2` is received (user-defined signal 2).

### User Input

### `Quit() <-chan struct{}`
Returns a channel that closes when the user types 'q' followed by Enter on stdin.

### `Enter() <-chan struct{}`
Returns a channel that closes when the user presses Enter.

### `Any() <-chan struct{}`
Returns a channel that closes on any key press.

### Composite Signals

### `AnyShutdown() <-chan struct{}`
Returns a channel that closes on SIGINT, SIGTERM, or user 'q'.

### `GracefulShutdown() <-chan struct{}`
Returns a channel that closes on SIGTERM or SIGHUP.

## Usage Examples

### Basic Signal Handling

```go
package main

import (
    "fmt"
    "time"
    
    "github.com/nyxstack/signals"
)

func main() {
    fmt.Println("Application started. Press Ctrl+C to exit.")
    
    // Wait for interrupt signal
    <-signals.Interrupt()
    
    fmt.Println("Received interrupt signal, shutting down...")
}
```

### Multiple Signal Handling

```go
package main

import (
    "fmt"
    "time"
    
    "github.com/nyxstack/signals"
)

func main() {
    fmt.Println("Server starting...")
    
    // Simulate some work
    go func() {
        for {
            fmt.Println("Working...")
            time.Sleep(2 * time.Second)
        }
    }()
    
    // Wait for any shutdown signal
    select {
    case <-signals.Interrupt():
        fmt.Println("Received SIGINT (Ctrl+C)")
    case <-signals.Terminate():
        fmt.Println("Received SIGTERM")
    case <-signals.Hangup():
        fmt.Println("Received SIGHUP")
    case <-signals.Quit():
        fmt.Println("User typed 'q'")
    }
    
    fmt.Println("Gracefully shutting down...")
}
```

### With Context for Graceful Shutdown

```go
package main

import (
    "context"
    "fmt"
    "time"
    
    "github.com/nyxstack/signals"
)

func main() {
    ctx, cancel := context.WithCancel(context.Background())
    defer cancel()
    
    // Start background work
    go worker(ctx)
    
    // Wait for shutdown signal
    select {
    case <-signals.Interrupt():
    case <-signals.Terminate():
    }
    
    fmt.Println("Shutdown signal received, stopping...")
    cancel()
    
    // Give time for graceful shutdown
    time.Sleep(1 * time.Second)
    fmt.Println("Application stopped")
}

func worker(ctx context.Context) {
    ticker := time.NewTicker(1 * time.Second)
    defer ticker.Stop()
    
    for {
        select {
        case <-ctx.Done():
            fmt.Println("Worker stopping...")
            return
        case <-ticker.C:
            fmt.Println("Working...")
        }
    }
}
```

## Use Cases

- **Web Servers**: Graceful shutdown on SIGTERM/SIGINT
- **CLI Tools**: Clean exit on Ctrl+C or user input
- **Kubernetes Applications**: Proper handling of pod termination signals
- **Daemon Processes**: Configuration reload on SIGHUP
- **Interactive Applications**: User-controlled exit with 'q' key

## Requirements

- Go 1.24.2 or later

## Contributing

Contributions are welcome! Please feel free to submit issues and pull requests. When contributing:

1. Follow Go conventions and best practices
2. Include tests for new signal handlers
3. Update documentation for any new functions
4. Ensure backward compatibility

## License

MIT License - see LICENSE file for details.