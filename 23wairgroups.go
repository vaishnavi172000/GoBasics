package main

import (
	"fmt"
	"sync"
	"time"
)

// worker function to complete the task
func worker1(i int) {
	fmt.Println("worker", i, "started")
	// added sleep to depict extensive task
	time.Sleep(time.Second)
	fmt.Println("worker", i, "done")
}

func main() {
	var wg sync.WaitGroup

	for i := 0; i < 3; i++ {
		wg.Add(1)

		// i should be captured in the closure
		// as clousure will be using this variable as its global variable
		// value ones updated is updated
		// as go routine will be called i will be increnetd to 3
		// hence add go routine will be called with 3
		go func(i int) {
			// defer should be called in same go routine which is waitgroup go routine
			// if decalared in different go routine then it will not work
			// it will cause deadlock
			defer wg.Done()
			worker1(i)
		}(i)
	}

	// will wait until all the go routines are completed
	wg.Wait()
}

// worker 2 started
// worker 0 started
// worker 1 started
// worker 1 done
// worker 2 done
// worker 0 done
