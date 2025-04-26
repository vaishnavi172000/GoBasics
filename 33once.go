package main

import (
	"fmt"
	"sync"
)

// once.Do() accepts a function as an argument.

// No matter how many times Do() is called (even from multiple goroutines),
// the function passed will be executed only once.
// Notes:
// If the function passed to Do() panics,
// it's considered as "executed" and will not run again.
// sync.Once is not resettable — once used, it can't be reused to call another fun

func worker2(i int) {
	fmt.Println("in the worker", i+1)

}

func main() {
	var once sync.Once
	var wg sync.WaitGroup

	for i := range 3 {
		wg.Add(1)
		go func(i int) {
			// once function will be executed once even it is called from the 3 go routines
			defer wg.Done()
			once.Do(func() {
				fmt.Println("in the Do once function", i+1)
			})
			// this function is executed in all go routines
			worker2(i)

		}(i)

	}

	// wait unteil exceution of all workers is done
	wg.Wait()

}

// vaishnavi@vaishnavis-MacBook-Air GoBasics % go run 33once.go
// in the Do once function 3
// in the worker 3
// in the worker 2
// in the worker 1
// vaishnavi@vaishnavis-MacBook-Air GoBasics %
