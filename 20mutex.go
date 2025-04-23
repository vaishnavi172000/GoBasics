package main

import (
	"fmt"
	"sync"
)

// mutex is used to avoid race condition
// we ahve seen atomic varible is used to avoid race condition
// but atomic varible is single counter
// to use bigger resource we have to use mutex
// mutex is used to avoid race condition

// conatiner conatins mutex ans map resource
type container struct {
	Mutex sync.Mutex
	m1    map[string]int
}

func main() {
	// define wait group
	var wg sync.WaitGroup

	c := &container{
		m1: map[string]int{},
	}

	names := []string{"one", "two", "three", "four", "five"}

	for i := 1; i <= 5; i++ {
		// add go routine
		wg.Add(1)
		e := i * 1000

		go func(i int, name string) {
			defer wg.Done()

			// lock the mutex before accessing the map
			c.Mutex.Lock()

			// unlock the mutex before leaving the go routine
			defer c.Mutex.Unlock()
			c.m1[name] = i

		}(e, names[i-1])
	}

	// wait for all the go routines to complete
	wg.Wait()

	fmt.Printf("\nValue of counter %+v\n", c.m1) // OUTPPUT :: Value of counter map[one:1000 two:2000 three:3000 four:4000 five:5000]
}
