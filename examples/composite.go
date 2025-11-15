package main

import (
	"fmt"

	"github.com/nyxstack/signals"
)

func main() {
	fmt.Println("Waiting for SIGINT, SIGTERM, or 'q'...")
	<-signals.AnyShutdown()
	fmt.Println("Shutdown signal received.")
}
