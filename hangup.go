package signals

import (
	"os"
	"os/signal"
	"syscall"
)

// Hangup returns a channel closed on SIGHUP (terminal hangup, reload, or session end).
func Hangup() <-chan struct{} {
	ch := make(chan struct{})
	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGHUP)
	go func() {
		defer signal.Stop(sigs)
		<-sigs
		close(ch)
	}()
	return ch
}
