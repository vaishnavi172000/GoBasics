package main

import (
	"fmt"
	"maps"
)

type s1 struct {
	var1 int
}

func main() {

	// map is decalredd as map[key data type]value_data_type
	var m1 map[string]int
	fmt.Println(m1, m1 == nil) // output: map[] true

	// this will create the panic situation as memory is not assigned to the map
	//m1["abc"] = 1

	// lets assign the memory to the map
	m1 = make(map[string]int)
	m1["abc"] = 1
	// Note that maps appear in the form map[k:v k:v] when printed with fmt.Println.
	fmt.Println(m1) // output :: map[abc:1]

	// every  element can be inserted in the array only once
	m1["abc"] = 2
	// above operaions oberode the avlues of key abc
	fmt.Println(m1) // output :: map[abc:2]

	// print the key for element which dose not exist
	fmt.Println(m1["pqr"]) // output :: 0

	// to check whether the key exist in the map or not
	if val, ok := m1["qwe"]; !ok { // output :: key qwe dose not exist in the map
		// val is value of key id not present it will give 0 value
		fmt.Println("key qwe dose nor exist in the map")
	} else {
		// if key is persent will get the value of key
		fmt.Println("value for key :", val)
	}

	if val, ok := m1["abc"]; !ok { // output :: value for key : 2
		fmt.Println("key qwe dose nor exist in the map")
	} else {
		fmt.Println("value for key :", val)
	}

	m1["abc"] = 0
	m1["pqr"] = 12
	// even this is having 0 as val but ok is true hence it will go in the else
	// hence tis appproach shiuld be used to check whether key is poresent or not in the map
	if val, ok := m1["abc"]; !ok { // output :: value for key : 0
		fmt.Println("key qwe dose nor exist in the map")
	} else {
		fmt.Println("value for key :", val)
	}

	// delete the element from thre map
	fmt.Println("Map before deleting the key abc", m1) //output :: Map before deleting the key abc map[abc:0 pqr:12]
	delete(m1, "abc")
	fmt.Println("Map after deleting key abc", m1) // output :: Map after deleting key abc map[pqr:12]

	// clear all the lements on the map
	// now map will not be null will conatin 0 elements
	clear(m1)
	fmt.Println(m1, m1 == nil, len(m1)) // output :: map[] false 0

	// can assign new keys to this cleared map
	m1["wer"] = 13

	// short hand assignment
	// can create map of kety elemnets which are comparable elements
	m2 := make(map[s1]string)

	b1 := s1{100}
	m2[b1] = "my world"
	fmt.Println(m2) // output :: map[{100}:my world]

	// use +v to print more precise
	fmt.Printf("%+v\n", m2) // output :: map[{var1:100}:my world]

	//decalration with directa assignment
	n := map[string]int{"foo": 1, "bar": 2}
	fmt.Println("map:", n) // output :: map: map[bar:2 foo:1]
	//The maps package contains a number of useful utility functions for maps.

	n2 := map[string]int{"foo": 1, "bar": 2} // output :: n==n2

	// maps are comparable entity in the go
	if maps.Equal(n, n2) {
		fmt.Println("n == n2")
	}

	// as underlying type of map is different comaprison od m1 and m2 is not allowed
	// if maps.Equal(m1,m2){

	// }

	// logic to just itterate over the map
	for k := range n {
		fmt.Println("key:", k)
	}
	//key: bar
	//key: foo
}
