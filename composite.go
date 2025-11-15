package signals

// AnyShutdown returns a channel closed on SIGINT, SIGTERM, or user 'q'.
func AnyShutdown() <-chan struct{} {
	ch := make(chan struct{})
	go func() {
		select {
		case <-Interrupt():
			close(ch)
		case <-Terminate():
			close(ch)
		case <-Quit():
			close(ch)
		}
	}()
	return ch
}

// GracefulShutdown returns a channel closed on SIGTERM or SIGHUP.
func GracefulShutdown() <-chan struct{} {
	ch := make(chan struct{})
	go func() {
		select {
		case <-Terminate():
			close(ch)
		case <-Hangup():
			close(ch)
		}
	}()
	return ch
}
