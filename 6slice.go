package main

import "fmt"

type book struct {
	pages int
}

// slice are fundamental notations in go
func main() {
	// slice are  more powefull interfaces than array in go
	// slice are dynamic in nature you can add elements to the slice

	// nil  slice decalration
	var a []int
	fmt.Println(a, a == nil) // output :: [] false
	// a[0] = 1 // this will create the panic as slice is nil
	// this is nil slice memory is not aasigned to it
	// hence this assignment create the panic
	// a[0] = 1

	// assign the memory to slice
	// this will create slice with len and capacity of 2
	a = make([]int, 2)
	a[0] = 1
	a[1] = 2
	fmt.Println(a) // output :: [1 2]
	// this will create the panic
	//a[2] = 3

	// append function returns the new slice
	a = append(a, 3)
	fmt.Println(a) // output :: [1 2 3]

	// appending the data to the slice more than 1 element
	a = append(a, 4, 5, 6)
	fmt.Println(a) // output :: [1 2 3 4 5 6]

	// appending the slice to the slice
	a = append(a, a...)
	fmt.Println(a) // output :: [1 2 3 4 5 6 1 2 3 4 5 6]

	// shorthand decalaration
	b := []string{"abc", "pqr"}
	fmt.Println(b) // output :: [abc pqr]

	// this assign memory and zero value to the slice elements
	c := make([]book, 2)
	fmt.Println(c) // output :: [{0} {0}]

	// 2d slice
	// here length will be 0 but the value will not be null
	d := [][]int{}
	fmt.Println(len(d), d == nil) // output :: 0 false
	d = make([][]int, 3)
	// this create the panic as length of the slice id zero
	d[0] = make([]int, 3)
	fmt.Println(d) // output :: [[0 0 0] [] []]

	// slice supports the slicing
	// here fist index is included and last index is excluded
	e := a[0:5]
	fmt.Println(e) // [1 2 3 4 5]

	// underlying slice is same
	// upadting the value of e  is result in upadtig the value of a
	e[0] = 100
	fmt.Println(a) //[100 2 3 4 5 6 1 2 3 4 5 6]
}
