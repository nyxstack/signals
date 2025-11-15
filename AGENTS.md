# AGENTS.md

## Project Overview

Simple channel-based signal handling package for Go. Each function returns a `<-chan struct{}` that closes when a specific signal or user input is received. Use in `select` statements to handle multiple signals concurrently.

## Setup Commands

- Build: `go build .`
- Test: `go test ./...`
- Install: `go get github.com/nyxstack/signals`

## Available Functions

### System Signals
- `signals.Interrupt()` - Closes on SIGINT (Ctrl+C). Use for immediate shutdown.
- `signals.Terminate()` - Closes on SIGTERM (Kubernetes/process managers). Use for graceful shutdown.
- `signals.Hangup()` - Closes on SIGHUP (terminal disconnect, config reload).
- `signals.Sigusr1()` - Closes on SIGUSR1 (Unix only). Custom application signal 1.
- `signals.Sigusr2()` - Closes on SIGUSR2 (Unix only). Custom application signal 2.

### User Input
- `signals.Quit()` - Closes when user types 'q' + Enter.
- `signals.Enter()` - Closes when user presses Enter.
- `signals.Any()` - Closes on any key press.

### Composite Signals
- `signals.AnyShutdown()` - Closes on SIGINT, SIGTERM, or 'q'. Use for universal shutdown handling.
- `signals.GracefulShutdown()` - Closes on SIGTERM or SIGHUP. Use for production graceful shutdown.

## Code Patterns

### Basic Usage
```go
<-signals.Interrupt()  // Blocks until SIGINT received
```

### Multiple Signal Handling
```go
select {
case <-signals.Interrupt():
    // Handle Ctrl+C
case <-signals.Terminate():
    // Handle SIGTERM
case <-signals.Hangup():
    // Reload configuration
}
```

### Background Goroutine Pattern
```go
go func() {
    <-signals.AnyShutdown()
    cleanup()
    os.Exit(0)
}()
```

### Web Server Shutdown
```go
go func() {
    if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
        log.Fatal(err)
    }
}()

<-signals.GracefulShutdown()
ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
defer cancel()
server.Shutdown(ctx)
```

## Platform Considerations

- `Sigusr1()` and `Sigusr2()` are Unix-only (build tag: `!windows`)
- All other functions work cross-platform
- Input functions (`Quit`, `Enter`, `Any`) read from stdin

## Code Style

- All functions return read-only channels: `<-chan struct{}`
- Channels close exactly once when signal/input received
- No panics, no resource leaks
- Each function spawns its own goroutine internally
- Safe for concurrent use

## Testing Instructions

When writing code that uses this package:
- Mock signal channels with your own `chan struct{}` in tests
- Close channels manually to simulate signal receipt
- Test timeout behavior with `time.After()` in select statements
- Verify cleanup code runs before process exit

## Common Use Cases

- **Development**: Use `Interrupt()` to quickly stop on Ctrl+C
- **Production**: Use `Terminate()` or `GracefulShutdown()` for container orchestration
- **Config Reload**: Use `Hangup()` for SIGHUP-based reload without restart
- **CLI Tools**: Use `Quit()` or `Any()` for interactive user control
- **Custom Logic**: Use `Sigusr1()`/`Sigusr2()` for application-specific triggers (Unix)
- **Universal Shutdown**: Use `AnyShutdown()` to handle multiple shutdown methods at once