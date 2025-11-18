package signals

// AnyShutdown returns a channel closed on SIGINT, SIGTERM, or user 'q'.
// The returned channel will close when any of these signals is received.
// Note: This function spawns goroutines for signal monitoring. Once the returned
// channel closes, those goroutines will be cleaned up automatically.
func AnyShutdown() <-chan struct{} {
	ch := make(chan struct{})
	interruptCh := Interrupt()
	terminateCh := Terminate()
	quitCh := Quit()

	go func() {
		select {
		case <-interruptCh:
			close(ch)
		case <-terminateCh:
			close(ch)
		case <-quitCh:
			close(ch)
		}
		// Note: The other signal goroutines will clean up when their channels
		// are garbage collected since no one is reading from them anymore.
	}()
	return ch
}

// GracefulShutdown returns a channel closed on SIGTERM or SIGHUP.
// The returned channel will close when any of these signals is received.
// Note: This function spawns goroutines for signal monitoring. Once the returned
// channel closes, those goroutines will be cleaned up automatically.
func GracefulShutdown() <-chan struct{} {
	ch := make(chan struct{})
	terminateCh := Terminate()
	hangupCh := Hangup()

	go func() {
		select {
		case <-terminateCh:
			close(ch)
		case <-hangupCh:
			close(ch)
		}
		// Note: The other signal goroutines will clean up when their channels
		// are garbage collected since no one is reading from them anymore.
	}()
	return ch
}
