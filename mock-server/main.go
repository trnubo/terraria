package main

import (
	"bufio"
	"fmt"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	// Set up a channel to handle signals
	signalChan := make(chan os.Signal, 1)
	signal.Notify(signalChan, syscall.SIGINT, syscall.SIGTERM)

	// Create a new scanner to read input from STDIN
	scanner := bufio.NewScanner(os.Stdin)

	for {
		select {
		case sig := <-signalChan:
			// Ignore SIGINT and SIGTERM signals
			fmt.Println("Received signal:", sig)
		default:
			if scanner.Scan() {
				input := scanner.Text()
				// Output the input with a '>' prefix
				fmt.Println(">", input)
				if input == "exit" {
					// Cleanly exit the application if "exit" is entered
					fmt.Println(">> Exiting the application.")
					return
				}
			}
		}
	}
}
