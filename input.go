package signals

import (
	"bufio"
	"os"
)

// Enter returns a channel closed when user presses Enter.
// The spawned goroutine will block on stdin.Read until Enter is pressed
// or stdin is closed. For applications that need explicit cleanup, consider
// using a context-based approach or ensuring stdin is properly closed.
func Enter() <-chan struct{} {
	ch := make(chan struct{})
	go func() {
		scanner := bufio.NewScanner(os.Stdin)
		if scanner.Scan() {
			close(ch)
		} else if err := scanner.Err(); err == nil {
			// stdin closed without error
			close(ch)
		}
	}()
	return ch
}

// Any returns a channel closed on any key press.
// The spawned goroutine will block on stdin.Read until any key is pressed
// or stdin is closed. For applications that need explicit cleanup, consider
// using a context-based approach or ensuring stdin is properly closed.
func Any() <-chan struct{} {
	ch := make(chan struct{})
	go func() {
		var b [1]byte
		n, err := os.Stdin.Read(b[:])
		if n > 0 || err == nil || err.Error() == "EOF" {
			close(ch)
		}
	}()
	return ch
}
