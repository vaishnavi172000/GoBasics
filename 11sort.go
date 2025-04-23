package main

import (
	"cmp"
	"fmt"
	"slices"
)

// in go slices package provides the sort function
// this sort method will sort the elements of the slice

// compareString will compare the length of the strings
func compareString(s1, s2 string) int {
	return cmp.Compare(len(s1), len(s2))
}

func main() {
	// slices.Sort will sort the elements of the slice
	nums := []int{6, 3, 1, 2, 5, 4}
	slices.Sort(nums)
	fmt.Println(nums) // OUTPUT :: [1 2 3 4 5 6]

	names := []string{"golang", "python", "java", "c++", "c"}
	slices.Sort(names)
	fmt.Println(names) // OUTPUT :: [c c++ golang java python]

	// as we provide the compareSrtring comparator function
	// this will sort the slice based on the length of the strings
	slices.SortFunc(names, compareString)
	fmt.Println(names) // OUTPUT :: [c c++ java golang python]

	type person struct {
		Name string
		Age  int
	}

	people := []person{
		{Name: "vaishnavi", Age: 24},
		{Name: "rasika", Age: 27},
		{Name: "vishal", Age: 20},
	}

	// this will sort the slice based on the age
	slices.SortFunc(people, func(p1, p2 person) int {
		return cmp.Compare(p1.Age, p2.Age)
	})
	fmt.Println(people) // OUTPUT :: [{vishal 20} {vaishnavi 24} {rasika 27}]
}
