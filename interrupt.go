package signals

import (
	"os"
	"os/signal"
	"syscall"
)

// Interrupt returns a channel closed on SIGINT (Ctrl+C).
func Interrupt() <-chan struct{} {
	ch := make(chan struct{})
	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGINT)
	go func() {
		defer signal.Stop(sigs)
		<-sigs
		close(ch)
	}()
	return ch
}
