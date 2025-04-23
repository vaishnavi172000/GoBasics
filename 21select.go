package main

import (
	"fmt"
	"time"
)

// select is a control structure
// that allows multiple channels to be read from or written
// to at the same time without blocking the thread
// select is a non blocking statement

func main() {
	mychannel1 := make(chan int)
	mychannel2 := make(chan int)

	// In your case, the select statement is trying to send values to mychannel1 and mychannel2,
	// but there is no other goroutine running concurrently that is receiving from these channels.
	// Therefore, the select statement deadlocks.
	select {
	case mychannel1 <- 1:
		fmt.Println("sent 1 to mychannel1")
	case mychannel2 <- 2:
		fmt.Println("sent 2 to mychannel2")
	default:
		fmt.Println("no channel was ready")
	} // output :: no channel was ready

	go func() {
		time.Sleep(1 * time.Second)
		mychannel1 <- 1
	}()
	go func() {
		time.Sleep(2 * time.Second)
		mychannel2 <- 2
	}()
	//We’ll use select to await both of these values simultaneously, printing each one as it arrives.

	for i := 0; i < 2; i++ {
		select {
		case msg1 := <-mychannel1:
			fmt.Println("received", msg1)
		case msg2 := <-mychannel2:
			fmt.Println("received", msg2)
		default:
			fmt.Println("no message received")
		}
	}
	// This is non blocker channel execution as it will  not stop
	// untill one of the conditon statement is true
	// output :: no message received
	// no message received

	for i := 0; i < 2; i++ {
		select {
		case msg1 := <-mychannel1:
			fmt.Println("received", msg1)
		case msg2 := <-mychannel2:
			fmt.Println("received", msg2)
		}
	}
	// output :: received 1
	// received 2
	// this is blocker channel execution
	// this is will stop untill one of statement is true

}
