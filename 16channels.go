package main

import "fmt"

func main() {
	// create a channel
	c := make(chan int)
	go func() {
		// send 42 to the channel
		c <- 42
	}()
	// this is syncrounius execution of the goroutine
	// this line will wait until the goroutine is executed
	// receive the value from the channel
	// hence execution will be blocked untill value in the channel is received
	fmt.Println(<-c)

	// create a buffered channel
	c2 := make(chan int, 2)
	c2 <- 1
	c2 <- 2
	// this does not error as the channel is buffered for size 2 only
	//c2 <- 3
	fmt.Println(<-c2)
	fmt.Println(<-c2)
	// this creates the panic as channel capacity is execceded
	//fmt.Println(<-c2)

	// create the done channel
	// This is the function we’ll run in a goroutine.
	// The done channel will be used to notify another goroutine that this function’s work is done.
	doneChannel := make(chan bool)
	go func() {
		fmt.Println("in function to test done cannel")
		for i := 0; i < 3; i++ {
			fmt.Println(i)
		}
		//Send a value to notify that we’re done.
		doneChannel <- true
	}()
	// Block until we receive a notification from the worker on the channel.
	fmt.Println("waiting till go routine send the done message in channle")
	// recieve the done value from the channel
	<-doneChannel
	// If you removed the <- done line from this program, the program would exit before the worker even started.

	//Closing a channel indicates that no more values will be sent on it.
	// This can be useful to communicate completion to the channel’s receivers.

}
