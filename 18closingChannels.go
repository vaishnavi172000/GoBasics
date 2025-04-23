package main

import "fmt"

// Closing a channel indicates that no more values will be sent on it.
// This can be useful to communicate completion to the channel’s receivers.

// In this example we’ll use a jobs channel to communicate work to be done from the main()
// goroutine to a worker goroutine.
// When we have no more jobs for the worker we’ll close the jobs channel.
func main() {
	jobs := make(chan int, 5)
	done := make(chan bool)

	go func() {
		for {
			j, more := <-jobs
			if more {
				fmt.Println("received job", j)
			} else {
				fmt.Println("received all jobs")
				done <- true
				return
			}
		}
	}()

	for j := 1; j <= 3; j++ {
		jobs <- j
		fmt.Println("sent job", j)
	}
	close(jobs)
	fmt.Println("sent all jobs")

	<-done

	_, ok := <-jobs
	fmt.Println("received more jobs:", ok)

	queue := make(chan string, 2)
	queue <- "one"
	queue <- "two"
	close(queue)
	// This range iterates over each element as it’s received from queue.
	// Because we closed the channel above, the iteration terminates after receiving the 2 elements.

	// This example also showed that it’s possible to close a non-empty channel
	// but still have the remaining values be received.
	for elem := range queue {
		fmt.Println(elem)
	}
}
