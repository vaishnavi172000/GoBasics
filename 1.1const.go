package main

import "fmt"

// In go we can define constants using the const keyword
// constants can be character, string, boolean or numeric values
// constants cannot be declared using the := syntax
// constants cannot be declared using the var keyword
// constants can be declared in a block
// constants can be declared in a block using the iota keyword
// iota is a predeclared identifier
// iota is used to declare successive integer constants
// iota starts from 0 and increments by 1 for each successive integer constant
// iota can be used to declare a block of constants
// iota can be used to declare a block of constants in a single line
func main() {
	// defining a constant of type float
	const Pi = 3.14
	fmt.Println("Pi:", Pi)

	// defining a constant of type string
	const Greeting = "Hello"
	fmt.Println("Greeting:", Greeting)

	// defining a block of constants
	const (
		// iota is used to declare successive integer constants
		// iota starts from 0 and increments by 1 for each successive integer constant
		// iota can be used to declare a block of constants
		// iota can be used to declare a block of constants in a single line
		First  = iota
		Second = iota
		Third  = iota
	)
	fmt.Println("First:", First)
	fmt.Println("Second:", Second)
	fmt.Println("Third:", Third)
}

// Pi: 3.14
// Greeting: Hello
// First: 0
// Second: 1
// Third: 2
