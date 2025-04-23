package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

// go supports atomic variables
// atomic variables are used to avoid race conditions
// atomic variable can avoid race condition for variable
// used by multiple resource that is go routine at a same time

func main() {
	var counter int64
	var c1 int64
	var wg sync.WaitGroup
	wg.Add(10)
	for i := 0; i < 10; i++ {
		go func() {
			defer wg.Done()
			for i := 0; i < 1000; i++ {
				// as counter is used by multiple go ruotine atomic will avoid race condition
				atomic.AddInt64(&counter, 1) // as race condition is not possible last value of counter will be 10000
			}
		}()
	}
	wg.Wait()

	fmt.Println("counter using atomic counter: ", counter) // counter using atomic counter:  10000

	var wg1 sync.WaitGroup
	wg1.Add(10)
	for i := 0; i < 10; i++ {
		go func() {
			defer wg1.Done()
			for i := 0; i < 1000; i++ {
				c1++ // as counter is used by multiple go ruotine will craete race condition
			}
		}()
	}
	wg1.Wait()

	fmt.Println("counter using normal counter: ", c1) // counter using normal counter:  9298

}
