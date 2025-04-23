package main

import "fmt"

// panic is a built in function
// panic is used to terminate the program
// everything after panic will not be executed

func main() {
	fmt.Println("before panic")
	panic("a problem")
	fmt.Println("after panic")

}

//output::
// [Running] go run "/Users/vaishnavi/Desktop/DSA/GoBasics/27panic.go"
// before panic
// panic: a problem

// goroutine 1 [running]:
// main.main()
// 	/Users/vaishnavi/Desktop/DSA/GoBasics/27panic.go:11 +0x59
// exit status 2

// [Done] exited with code=1 in 0.707 seconds
