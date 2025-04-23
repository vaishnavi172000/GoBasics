package main

// go support basic  values like  int string float bool

import "fmt"

func main() {
	// in go strings can be added usinh + operator
	fmt.Println("go" + "lang")
	// go support integer and float values
	fmt.Println("1+1=", 1+1)
	fmt.Println("7.0/3.0=", 7.0/3.0)
	// go support boolean values
	fmt.Println(true && false)
	fmt.Println(true || false)
	fmt.Println(!true)
}
