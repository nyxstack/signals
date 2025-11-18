package signals

import (
	"bufio"
	"os"
)

// Quit returns a channel closed when user types 'q' + Enter.
// The spawned goroutine will continue reading stdin until 'q' is entered
// or stdin is closed. For applications that need explicit cleanup, consider
// using a context-based approach or ensuring stdin is properly closed.
func Quit() <-chan struct{} {
	ch := make(chan struct{})
	go func() {
		scanner := bufio.NewScanner(os.Stdin)
		for scanner.Scan() {
			if scanner.Text() == "q" {
				close(ch)
				return
			}
		}
		// If scanner stops (e.g., stdin closed), close the channel
		if err := scanner.Err(); err == nil {
			close(ch)
		}
	}()
	return ch
}
