package main

import "fmt"

// In go we ca define variable any where in the code
// but it is a good practice to define all the variables at the beginning of the function
// this will help in understanding the code better
// go is a statically typed language
// this means that once a variable is declared it cannot be changed to another type
// go will automatically detect the type of variable based on the value assigned
// go will assign zero values as values to variables if not assigned
// zero values are 0 for numeric types, false for boolean types and "" for strings
// go also supports short hand assignment
// go also supports multiple variable declaration in a single line
// go also supports declaring variables without mentioning the type
// go also supports declaring variables without mentioning the type and assigning values
func main() {
	// decalring a variable
	// here type is not mentioned so go will automatically detect the type of variable
	var a = "initial"
	fmt.Println(a) // output :: initial

	// multiple variables can be declared in a single line
	// type of variables can be mentioned explicitly
	var b, c int = 1, 2
	fmt.Println(b, c) // output :: 1 2

	// go will automatically detect the type of variable
	// here type is detected as float based on the value assigned
	// this is short hand assignment
	d := 12.67
	fmt.Println(d) // output :: 12.67

	// go will assign zero values as values to variables if not assigned
	// zero values are 0 for numeric types, false for boolean types and "" for strings
	var a1 int
	fmt.Println(a1) // output :: 0
}
