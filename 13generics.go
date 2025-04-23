package main

import (
	"fmt"
)

// creating generic function using any
// all data type slice can be passed to this fucntion now
func printSlice[T any](s []T) {
	for i, v := range s {
		fmt.Println("indesx : ", i, "value:", v)
	}
}

// creating generic function using interface
// all data type slice can be passed to this fucntion now
func printSlice1[T interface{}](s []T) {
	for i, v := range s {
		fmt.Println("indesx : ", i, "value:", v)
	}
}

// creating generic function using interface
// all data type slice can be passed to this fucntion now
// comparable can use to allow all comparable types
// 2 diffrent type swe can pass
func printSlice3[T comparable, V any](s []T, m V) {
	for i, v := range s {
		fmt.Println("indesx : ", i, "value:", v)
	}
	fmt.Println(m)
}

// creating generic function
//
//	int and string slice can be passed to this fucntion now
func printSlice2[T int | string](s []T) {
	for i, v := range s {
		fmt.Println("indesx : ", i, "value:", v)
	}
}

type stack[T any] struct {
	elements []T
}

func main() {
	nums := []int{1, 2, 3}
	names := []string{"At", "Vaishnavi"}

	printSlice(nums)
	printSlice(names)

	printSlice1(nums)
	printSlice1(names)

	printSlice2(nums)
	printSlice2(names)

	printSlice3(nums, 1)
	printSlice3(names, "cbdsf")

	st := stack[string]{
		elements: []string{"golang"},
	}

	fmt.Println(st)        // {[golang]}
	fmt.Printf("%+v ", st) // {elements:[golang]}
}
