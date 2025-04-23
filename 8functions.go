package main

import "fmt"

// In go functions are defined using the func keyword
// A function can take zero or more arguments
// A function can return zero or more values
// A function is a type in go
// A function can be passed as an argument to another function
// A function can be returned as a value from another function
// A function can be assigned to a variable
// A function can be used as a type
// A function can be used as a type in a struct
// A function can be used as a type in an interface
// A function is go supports recursion

// defining a function that takes two arguments and returns a value
func add(a int, b int) int {
	return a + b
}

// defining a function that takes three arguments and returns a value
// here if function argumnet types are same then we can mention the type only once
// a, b, c int
// but we cannot define it like this( a int, b, c int) this is not allowed in go
// we can also return the value without mentioning the variable name a
func add3nums(a, b, c int) (d int) {
	d = a + b + c
	return
}

// defining a function that takes two arguments and returns two values
// go supports multiple return values
func swap(a, b int) (int, int) {
	return b, a
}

// VARIDIC FUNCTIONS
// defining a function that takes a variable number of arguments
// go supports variadic functions
// variadic functions can be called with any number of trailing arguments
// variadic functions can be called with zero arguments
func sum(nums ...int) int {
	total := 0
	for _, num := range nums {
		total += num
	}
	return total
}

// ANONYMOUS FUNCTIONS
// defining an anonymous function
// anonymous functions are functions without a name
// anonymous functions are defined using the func keyword followed by the arguments and the code block
// anonymous functions can be assigned to a variable
// anonymous functions can be passed as an argument to another function
// anonymous functions can be returned as a value from another function
// anonymous functions can be used as a type
// anonymous functions can be used as a type in a struct
// anonymous functions can be used as a type in an interface
func anonymous() {

	// This is an anonymous function
	func() {
		fmt.Println("anonymous function")
	}()
}

// defining a function that returns an anonymous function
func returnanonymous() func() {
	return func() {
		fmt.Println("returning anonymous function")
	}
}

// defining a function that takes a  anonymous function as an argument
func callanonymous(f func(int, int) int) int {
	d := f(1, 2)
	return d
}

// A closure is a special type of anonymous function that references variables declared outside of the function itself.
// It is similar to accessing global variables which are available before the declaration of the function.
func closure() func() int {
	x := 0
	return func() int {
		x++
		return x
	}
}

// PASS BY VALUE
// go supports pass by value
// when we pass a variable to a function it is passed by value
// this means that a copy of the variable is passed to the function
// hence any changes made to the variable inside the function will not reflect outside the function
func zeroVal(i int) {
	i = 10000

}

// PASS BY REFERENCE
// go supports pass by reference
// when we pass a pointer to a variable to a function it is passed by reference
// this means that the address of the variable is passed to the function
// hence any changes made to the variable inside the function will reflect outside the function
func zeroPtr(i *int) {
	*i = 10000
}

func main() {
	// calling the function
	fmt.Println(add(1, 2)) // output :: 3

	// calling varidic function with single argument
	fmt.Println("sum of single digit ", sum(1)) // output :: sum of single digit  1

	//calling varidic function with multiple arguments
	fmt.Println("sum of multiple digits ", sum(1, 2, 3, 4, 5)) // output ::  sum of multiple digits  15

	// calling the varidic function with slice this will add ... at the end
	// in veridic function varidic argument is converted to slice
	// hence passins slice will be added as varidc variable
	num := []int{1, 2, 3, 56}
	fmt.Println("sum of slice elements", sum(num...)) // output :: sum of slice elements 62

	// calling the anonymous function
	anonymous() // output :: anonymous function

	// we can assign the anonymous function to a variable
	f := returnanonymous()
	// calling the anonymous function using the variable
	f() // output :: returning anonymous function

	// calling the function that takes anonymous function as an argument
	// here we are passing the anonymous function as an argument
	// this is also called as callback function
	// this is used in the callback functions
	fmt.Println("calling the function that takes anonymous function as an argumentL ", callanonymous(func(a, b int) int {
		return a + b
	})) // output :: calling the function that takes anonymous function as an argument 3

	// calling the closure function
	// here variable x is declared outside the function
	// it acts as a global variable
	// hence the value of x is retained between the function calls
	// but this value is restriced to one variable only
	// hence the value of x is retained between the function calls
	// but this value is restriced to one variable only
	f1 := closure()
	fmt.Println("closure function", f1()) // output :: closure function 1
	fmt.Println("closure function", f1()) // output :: closure function 2
	fmt.Println("closure function", f1()) // output :: closure function 3

	// calling the closure function
	// here variable x is redeclared outside the function
	// it acts as a global variable
	// value of x will start again from 0
	f2 := closure()
	fmt.Println("closure function", f2()) // output :: closure function 1
	fmt.Println("closure function", f2()) // output :: closure function 2
	// x avlue for f2 is different from f1

	// calling the pass by value function
	// here the value of i is passed to the function
	i := 10
	zeroVal(i)
	fmt.Println("pass by value", i) // output :: pass by value 10

	// calling the pass by reference function
	// here the address of i is passed to the function
	zeroPtr(&i)
	fmt.Println("pass by reference", i) // output :: pass by reference 10000

}
