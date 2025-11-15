//go:build !windows

package signals

import (
	"os"
	"os/signal"
	"syscall"
)

// Sigusr2 returns a channel closed on SIGUSR2 (user-defined signal 2).
func Sigusr2() <-chan struct{} {
	ch := make(chan struct{})
	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGUSR2)
	go func() {
		defer signal.Stop(sigs)
		<-sigs
		close(ch)
	}()
	return ch
}
