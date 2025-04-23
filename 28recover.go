package main

import "fmt"

// recover is a built in function
// recover is used to recover from a panic
// recover is  always used inside a defer
func main() {
	defer func() {

		if r := recover(); r != nil {
			fmt.Println("Recovered in f", r)
		}

	}()

	fmt.Println("before panic")

	panic("a problem")
}

// before panic
// Recovered in f a problem
