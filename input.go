package signals

import (
	"bufio"
	"os"
)

// Enter returns a channel closed when user presses Enter.
func Enter() <-chan struct{} {
	ch := make(chan struct{})
	go func() {
		scanner := bufio.NewScanner(os.Stdin)
		if scanner.Scan() {
			close(ch)
		}
	}()
	return ch
}

// Any returns a channel closed on any key press.
func Any() <-chan struct{} {
	ch := make(chan struct{})
	go func() {
		var b [1]byte
		os.Stdin.Read(b[:])
		close(ch)
	}()
	return ch
}
