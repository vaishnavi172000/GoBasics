package main

import (
	"errors"
	"fmt"
)

// Go’s structs are typed collections of fields.
// They’re useful for grouping data together to form records.

type person struct {
	name string
	age  int
}

// new Person constructs a new person struct with the given name.

//  is a garbage collected language; you can safely return a pointer to a local variable -
//  it will only be cleaned up by the garbage collector when there are no active references to it.

func newPerson(name string, age int) (*person, error) {
	// here we can add validation on the input provided
	if age < 1 {
		errors.New("please provide positive age")
	}
	p := person{name: name, age: age}
	return &p, nil
}

// embedded structs
// Go supports embedding of structs and interfaces to express a more seamless composition of types
type employee struct {
	// An embedding looks like a field without a name.
	person
	salary int
}

func main() {

	// deaclaring the person variable
	p := person{age: 10, name: "rupali"}
	fmt.Println(p)       // :: output {rupali 10}
	fmt.Printf("%+v", p) // output :: {name:rupali age:10}

	// structure menbers are accesible using th dot operator
	fmt.Println("name : ", p.name)

	// Anonomous struct
	// If a struct type is only used for a single value,
	// we don’t have to give it a name. The value can have an anonymous struc
	dog := struct {
		name   string
		isGood bool
	}{
		"Rex",
		true,
	}
	fmt.Println(dog)        // output :: {Rex true}
	fmt.Printf("%T\n", dog) // output :: struct { name string; isGood bool }

	emp := employee{
		person: person{
			name: "vaishnavi",
			age:  24,
		},
		salary: 100000,
	}
	fmt.Printf("\n%+v\n", emp) // output :: {person:{name:vaishnavi age:24} salary:100000}

	fmt.Println(emp.name)        // can access directly
	fmt.Println(emp.person.name) // can access via the person element also

}
