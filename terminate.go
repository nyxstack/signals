package signals

import (
	"os"
	"os/signal"
	"syscall"
)

// Terminate returns a channel closed on SIGTERM (used by Kubernetes).
func Terminate() <-chan struct{} {
	ch := make(chan struct{})
	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGTERM)
	go func() {
		defer signal.Stop(sigs)
		<-sigs
		close(ch)
	}()
	return ch
}
