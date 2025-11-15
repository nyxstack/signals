//go:build !windows

package signals

import (
	"os"
	"os/signal"
	"syscall"
)

// Sigusr1 returns a channel closed on SIGUSR1 (user-defined signal 1).
func Sigusr1() <-chan struct{} {
	ch := make(chan struct{})
	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGUSR1)
	go func() {
		defer signal.Stop(sigs)
		<-sigs
		close(ch)
	}()
	return ch
}
