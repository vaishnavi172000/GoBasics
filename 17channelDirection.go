package main

import "fmt"

// when we pass the channel to the function it is reciver or sender based on the direction of the channel
// if the channel is reciver then is is defined as (name <-channel)
// id the channel is sender then it is defined as (name channel <-)
// it is not allowed to read the value from the channel which is sender
// it is not allowed to write the value to the channel which is reciver

// ping will recieve the ping message and send to the channel
// hence channel type here is the sender channel
func ping(c chan<- string, message string) {
	c <- message
}

// pong will recieve the ping message and send to the channel
// hence c channel type here is the reciver channel
// this function is not allowed to write the value to the channel
// it is allowed to read the value from the channel
// this is wrire message in message seder channel
func pong(c <-chan string, message chan<- string) {
	m1 := <-c
	message <- m1
}

func main() {

	// declare the channel of type string
	pings := make(chan string)
	pongs := make(chan string)

	// here ping is passed as sender
	go ping(pings, "passed message")
	// here pong is passed as reciver
	go pong(pings, pongs)

	msg := <-pongs
	fmt.Println(msg)

}
