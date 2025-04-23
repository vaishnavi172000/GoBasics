package main

import "fmt"

func main() {
	// go supports for loop
	// for loop can be used to iterate over a slice, array, string or map
	// for loop can be used to iterate over a slice, array, string or map using range keyword
	// range keyword returns two values
	// first value is the index and second value is the value at that index

	// normal for loop syntax
	sum := 0
	for i := 0; i < 10; i++ {
		sum += i
	}
	fmt.Println("printing the sum cacalutaled using simple for loop for 0 to 9", sum)
	// output :: printing the sum cacalutaled using simple for loop for 0 to 9 45

	// for loop to iterate over a slice
	nums := []int{2, 3, 4}
	sum = 0
	for i := 0; i < len(nums); i++ {
		sum += nums[i]
	}
	fmt.Println("printing the sum cacluated of the slice using simple for loop", sum)
	// output :: printing the sum cacluated of the slice using simple for loop 9

	// for loop to iterate over a slice using range keyword
	sum = 0
	for i, num := range nums {
		sum += num
		fmt.Println(i, num)
		//output :: 0 2
		// 1 3
		// 2 4
	}
	fmt.Println("printing the sum cacluated using range for loop", sum)
	// output :: printing the sum cacluated using range for loop 9

	// Creating a Map in Go
	var ages map[string]int

	// Initializing the Map
	ages = make(map[string]int)

	ages["alice"] = 31
	ages["charlie"] = 34
	fmt.Println("printing the map using for loop")

	for name, age := range ages {
		fmt.Println(name, age)
	}
	// output :: alice 31
	// charlie 34

	// for loop to iterate over a string
	// range keyword can be used to iterate over a string
	fmt.Println("printing the string using for loop")
	str := "vaishnavi"
	for i, c := range str {
		fmt.Println(i, string(c))
	}
	// output :: 0 v
	// 1 a
	// 2 i
	// 3 s
	// 4 h
	// 5 n
	// 6 a
	// 7 v
	// 8 i

	// for loop to iterate over a string
	// range keyword can be used to iterate over a string
	// if we want to ignore the index we can use _ instead of i
	println("printing the string using for loop ignoring the indesx")
	for _, c := range str {
		fmt.Print(string(c))
	}

	// infinite loop
	// for {
	// 	fmt.Println("infinite loop")
	// }

	fmt.Println("printing the numbers using for loop using range just for the index")
	// for loop to interate over the number using range
	for i := range nums {
		fmt.Println(i)
	}
	// output :: 0
	// 1
	// 2

}
