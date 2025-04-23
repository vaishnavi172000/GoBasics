package main

import "fmt"

func main() {

	// arrays in go are fixed size
	// arrays in go are of type value
	// arrays in go are of type value and not reference

	// declare a int aray of size 5
	// by default all the elements in the array are set to 0
	var a [5]int
	fmt.Println("print the arrya with deafult values", a)
	// output :: print the arrya with deafult values [0 0 0 0 0]

	// decalre a int array of size 5 and assign values to the array
	// go also supports short hand assignment
	b := [5]int{1, 2, 3, 4, 5}
	fmt.Println("print the array with values assigned", b)
	// output :: print the array with values assigned [1 2 3 4 5]

	// declare a int array of size 5 and assign values to the array
	var c = [5]int{1, 2, 3, 4, 5}
	fmt.Println("print the array with values assigned", c)
	// output :: print the array with values assigned [1 2 3 4 5]

	// decalrethe int array where length is infrred from the number of values assigned
	d := [...]int{1, 2, 3, 4, 5}
	fmt.Println("print the array with values assigned", d)
	// output :: print the array with values assigned [1 2 3 4 5]

	// itterating over thre array
	// using the normal for loop
	for i := 0; i < len(b); i++ {

		fmt.Println("index:", i, "value:", b[i])
	}
	// output :: index: 0 value: 1
	// index: 1 value: 2
	// index: 2 value: 3
	// index: 3 value: 4
	// index: 4 value: 5

	// upadting the value of the array
	b[2] = 100
	fmt.Println("print the array after updating the value", b)
	// itterating over thre array
	// using the range keyword
	for i, v := range b {
		fmt.Println("index:", i, "value:", v)
	}
	// output :: index: 0 value: 1
	// index: 1 value: 2
	// index: 2 value: 100
	// index: 3 value: 4
	// index: 4 value: 5

	// iterating over the array
	// just for index
	for i := range b {
		fmt.Println("index:", i)
	}
	// output :: index: 0
	// index: 1
	// index: 2
	// index: 3
	// index: 4

	// 2d array in go
	// declare a 2d array
	var twoD [2][3]int
	fmt.Println("print the 2d array with default values", twoD)
	// output :: print the 2d array with default values [[0 0 0] [0 0 0]]

	two01 := [2][3]int{{1, 2, 3}, {4, 5, 6}}
	fmt.Println("print the 2d array with values assigned", two01)
	// output :: print the 2d array with values assigned [[1 2 3] [4 5 6]]

}

// output::
// print the arrya with deafult values [0 0 0 0 0]
// print the array with values assigned [1 2 3 4 5]
// print the array with values assigned [1 2 3 4 5]
// print the array with values assigned [1 2 3 4 5]
// index: 0 value: 1
// index: 1 value: 2
// index: 2 value: 3
// index: 3 value: 4
// index: 4 value: 5
// print the array after updating the value [1 2 100 4 5]
// index: 0 value: 1
// index: 1 value: 2
// index: 2 value: 100
// index: 3 value: 4
// index: 4 value: 5
// index: 0
// index: 1
// index: 2
// index: 3
// index: 4
// print the 2d array with default values [[0 0 0] [0 0 0]]
// print the 2d array with values assigned [[1 2 3] [4 5 6]]
