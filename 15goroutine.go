package main

import (
	"fmt"
	"time"
)

// A goroutine is a lightweight thread of execution.

func f1(str string) {
	for i, r := range str {
		fmt.Println("index: ", i, "value", r)
	}
}

func main() {

	// normla synronous floe of the fucntion
	f1("abc")

	go f1("pqr")

	go func(str string) {
		fmt.Println("in anonnymous function on routine")
		fmt.Println(str)
	}("vaishnavi")

	time.Sleep(time.Second)

	fmt.Println("Done")

}

//index:  0 value 97
// index:  1 value 98
// index:  2 value 99
// in anonnymous function on routine
// vaishnavi
// index:  0 value 112
// index:  1 value 113
// index:  2 value 114
// Done
