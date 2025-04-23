package main

import (
	"fmt"
	"time"
)

// write a code to spawn three workes for three
// task 5
// no do not juse any advange topic just go ruotine and channels

// here worker is the goroutine
// now we can define the number of goroutine we want
// leading to the concurrency
func worker(id int, jobs <-chan int, results chan<- int) {
	// will wait till jobs channel is not closed
	for j := range jobs {
		fmt.Println("Starting job with worker: ", id, "job id:", j)
		time.Sleep(time.Second)
		results <- j * 2
		fmt.Println("Completed job with worker: ", id, "job id:", j)
	}

}

func main() {
	// here number of jobs is 5
	jobs := make(chan int, 5)
	resultd := make(chan int, 5)

	// number of worker threads required
	noOfWorkers := 3

	// start the workers goroutines
	for i := 0; i < noOfWorkers; i++ {
		go worker(i, jobs, resultd)
	}

	// pass value to the channel
	for i := 0; i < 5; i++ {
		jobs <- i + 1
	}

	// close the jobs channel to indicate the work is done
	close(jobs)

	// collect the results

	// this can also be used
	// for i := 0; i < 5; i++ {
	// 	result := <-resultd
	// 	fmt.Println("Result: for job id:", i, "result:", result)
	// }

	i := 0
	for r := range resultd {
		fmt.Println("Result: for job id:", i, "result:", r)
		i++
		if i == 5 {
			// need to close else for loop will run forever and block when othing is left to read
			close(resultd)
		}

	}
}

// here we can see that 2 workers are reused
// we ahd 5 jobd so equal to 5 goroutines
// but we only have 3 workers
// so only 3 goroutines are used
//  this is will 2 sec takes if not done with go rotines would have taken 5 seconds

// OUTPUT

// Starting job with worker:  2 job id: 1
// Starting job with worker:  0 job id: 2
// Starting job with worker:  1 job id: 3
// Completed job with worker:  1 job id: 3
// Starting job with worker:  1 job id: 4
// Completed job with worker:  2 job id: 1
// Starting job with worker:  2 job id: 5
// Completed job with worker:  0 job id: 2
// Result: for job id: 0 result: 6
// Result: for job id: 1 result: 2
// Result: for job id: 2 result: 4
// Completed job with worker:  2 job id: 5
// Result: for job id: 3 result: 10
// Completed job with worker:  1 job id: 4
// Result: for job id: 4 result: 8
