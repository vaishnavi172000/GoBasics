package main

import "fmt"

// repanic is a suitation when you want to recover from a panic and then panic again
// this us used to log the error and then panic again
// this will avoid silent panic

func main() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in f", r)
		}
		panic("a bigger problem")
	}()
	fmt.Println("before panic")
	panic("a problem")
}

// [Running] go run "/Users/vaishnavi/Desktop/DSA/GoBasics/29repanic.go"
// before panic
// Recovered in f a problem
// panic: a problem [recovered]
// 	panic: a bigger problem

// goroutine 1 [running]:
// main.main.func1()
// 	/Users/vaishnavi/Desktop/DSA/GoBasics/29repanic.go:10 +0x30
// panic({0x152dd40?, 0x1547e00?})
// 	/usr/local/go/src/runtime/panic.go:792 +0x132
// main.main()
// 	/Users/vaishnavi/Desktop/DSA/GoBasics/29repanic.go:13 +0x78
// exit status 2

// [Done] exited with code=1 in 0.72 seconds
