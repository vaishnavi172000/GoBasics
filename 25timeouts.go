package main

import (
	"fmt"
	"time"
)

// main demonstrates the use of timeouts in channels.
//
// Timeouts are used when we're not sure whether a goroutine will send data to
// a channel or not. If no data is sent to the channel within the specified
// time, the timeout is triggered and the code in the timeout case block is
// executed.
//
// The first select statement will print "Value recieved from channel : 10"
// because the goroutine will send data to the channel within the 5 second
// timeout period.
//
// The second select statement will print "time out for reciveing value from
// channel" because the goroutine will not send data to the channel within the
// 1 second timeout period.
func main() {
	ch := make(chan int)

	go func() {
		fmt.Println("in go routine function")
		time.Sleep(time.Second * 5)
		ch <- 10
	}()

	var m1 int

	select {
	case m1 = <-ch:
		{
			fmt.Println("Value recieved from channel :", m1)
		}
	case <-time.After(time.Second * 5):
		{
			fmt.Println("time out for reciveing value from channel")
		}
	}

	select {
	case m1 = <-ch:
		{
			fmt.Println("Value recieved from channel :", m1)
		}
	case <-time.After(time.Second * 1):
		{
			fmt.Println("time out for reciveing value from channel")
		}
	}

}

// output::
// in go routine function
// time out for reciveing value from channel
// Value recieved from channel : 10
