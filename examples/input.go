package main

import (
	"fmt"
	"github.com/nyxstack/signals"
)

func main() {
	fmt.Println("Press Enter to continue...")
	<-signals.Enter()
	fmt.Println("Enter pressed!")

	fmt.Println("Press any key to continue...")
	<-signals.Any()
	fmt.Println("Any key pressed!")

	fmt.Println("Type 'q' and press Enter to quit...")
	<-signals.Quit()
	fmt.Println("Quit command received!")
