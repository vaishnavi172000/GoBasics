
package main

import (
	"fmt"
	"time"
)

// ticker is a functionality in go
// this is allow you to do the task regularly after specific time interval
// Main function
func main() {
	// Create a new ticker that triggers every 5 seconds
	tc := time.NewTicker(time.Second * 5)

	// Create an unbuffered channel to signal the goroutine to exit
	done := make(chan bool)

	// Start a new goroutine that runs concurrently with the main goroutine
	go func() {
		// Loop indefinitely until the goroutine is signaled to exit
		for {
			// Use a select statement to wait for either a tick from the ticker or a signal on the done channel
			select {
			case t := <-tc.C:
				// If a tick is received, print the current time
				fmt.Println("tick at : ", t)
			case <-done:
				// If a signal is received on the done channel, print a message and exit the goroutine
				fmt.Println("exiting go routine")
				return
			}
		}
	}()

	// Sleep for 17 seconds to allow the ticker to trigger a few times
	time.Sleep(time.Second * 17)

	// Stop the ticker to prevent further triggers
	tc.Stop()

	// Signal the goroutine to exit by sending a value on the done channel
	done <- true

}
