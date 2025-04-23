package main

import "fmt"

// Go supports methods defined on struct types.
// Methods can be defined for either pointer or value receiver types.

// define base struct to define the reciver methods
type rect struct {
	width  int
	height int
}

// create the new instance of the rectangle object
func NewRectangle(width, height int) *rect {
	r := rect{
		width:  width,
		height: height,
	}
	return &r
}

// this is reciver function of rect and
// this is defined on pointer reciever variable
func (rt *rect) Area() int {
	return rt.height * rt.width
}

// this is reciver function of rect
// this is defined on value reciver variable
func (rt rect) Perimeter() int {
	return 2 * (rt.height + rt.width)
}

// this is not upadte the value of actual varaible
func (rt rect) Update() {
	rt.height = 100
}

// this will upadte the value of actual variable
func (rt *rect) updatepr() {
	rt.height = 100
}

func main() {

	// create new rect variable
	rt := NewRectangle(1, 2)
	fmt.Println(rt.Area())      // output :: 2
	fmt.Println(rt.Perimeter()) // output :: 6

	rt1 := *NewRectangle(4, 5)
	fmt.Println(rt1.Area())      // output :: 20
	fmt.Println(rt1.Perimeter()) // output :: 18
	rt1.Update()
	fmt.Println(rt1) // output :: {4 5}
	rt1.updatepr()
	fmt.Println(rt1) // output :: {4 100}

}
