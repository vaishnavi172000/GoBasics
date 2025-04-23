package main

import (
	"fmt"
	"time"
)

// we use the ticker when we want to do some task in the future after specificc time period
// time.Sleep() also wait for time but diffrence is if we want we can stop the ticker before ots time to fire

// Main function where the program starts execution
func main() {
	// Create a new timer that will expire after 2 seconds
	timer := time.NewTimer(2 * time.Second)

	// Create a new timer that will expire after 5 seconds
	timer1 := time.NewTimer(5 * time.Second)

	// Start a new goroutine that waits for timer1 to expire
	go func() {
		// Wait for the timer to expire, then print a message
		<-timer1.C
		fmt.Println("Timer for 5 seconds is fired")
	}()

	// Start a new goroutine that waits for timer to expire
	go func() {
		// Wait for the timer to expire, then print a message
		<-timer.C
		fmt.Println("Timer for 2 seconds is fired")
	}()

	// Pause the main goroutine for 3 seconds
	time.Sleep(time.Second * 3)

	// Stop timer1, preventing it from expiring
	timer1.Stop()
	fmt.Println("Timer for 5 seconds is stopped")

	// Pause the main goroutine for an additional 5 seconds
	time.Sleep(time.Second * 5)
}
