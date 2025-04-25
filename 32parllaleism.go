package main

import (
	"fmt"
	"runtime"
	"time"
)

// In Go, goroutines are light-weight threads of execution.
// To run this routines in parallel we can use go keyword
// to run in parallel not concurrent we will run go routines in diffrent cpu cores
// This is called parallelism
// for this we ahve set GOMAXPROCS to use all the cores available

func compute(id int) {
	start := time.Now()
	count := 0
	for i := 0; i < 1e8; i++ {
		count += i
	}
	fmt.Printf("Goroutine %d done in %v\n", id, time.Since(start))
}

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU()) // use all available CPUs

	go compute(1)
	go compute(2)

	time.Sleep(3 * time.Second) // Wait for both
}
