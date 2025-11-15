package main

import (
	"fmt"

	"github.com/nyxstack/signals"
)

func main() {
	fmt.Println("Waiting for Ctrl+C (SIGINT)...")
	<-signals.Interrupt()
	fmt.Println("Received SIGINT, exiting.")
}
