package main

import "fmt"

func main() {

	// go supports if else statements
	// parentheses are not required around the condition in go
	// braces are required around the body
	if 8%2 == 0 || 7%2 == 0 {
		fmt.Println("either 8 or 7 are even")
	}
	//A statement can precede conditionals;
	//any variables declared in this statement are available in the current and all subsequent branches.

	if num := 9; num < 0 {
		fmt.Println(num, "is negative")
	} else if num < 10 {
		fmt.Println(num, "has 1 digit")
	} else {
		fmt.Println(num, "has multiple digits")
	}

	// go does not support ternary if else operator
}
