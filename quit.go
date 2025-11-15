package signals

import (
	"bufio"
	"os"
)

// Quit returns a channel closed when user types 'q' + Enter.
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
	}()
	return ch
}
