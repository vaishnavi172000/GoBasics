package main

import "fmt"

// init function is called before main function
// when multiple init functions are present in the same package,
// they are executed in the order they are defined in the source code

func init() {
	fmt.Println("First init")
}

func init() {
	fmt.Println("Second init")
}

func init() {
	fmt.Println("Third init")
}

func init() {
	fmt.Println("Fourth init")
}

func main() {}

// Output:
// First init
// Second init
// Third init
// Fourth init
