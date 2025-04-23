package main

import "fmt"

// Interfaces are named collections of method signatures.
type shape interface {
	area() int
	perimeter() int
}

// we will implement shape interface for both circle and rectangle
type circle struct {
	radius int
}

func (ct circle) area() int {
	return int(3.14 * float64(ct.radius) * float64(ct.radius))
}

func (ct circle) perimeter() int {
	return int(2 * 3.14 * float64(ct.radius))
}

type rectangle struct {
	height int
	width  int
}

func (rt rectangle) area() int {
	return rt.height * rt.width
}

func (rt rectangle) perimeter() int {
	return 2 * rt.height * rt.width
}

// we passed circle and rectangle varibles as argumnets
// this structs are auto converted to the interface
func display(s shape) {
	// typecast the shape to circle
	if c, ok := s.(circle); ok {
		fmt.Println("This shape is circle of radius :", c.radius)
	}
	// type case the shape to rectagle
	if rt, ok := s.(rectangle); ok {
		fmt.Println("This shape is rectangke with the width : ", rt.width, "and height :", rt.height)
	}
	// interfcae fuction will call base call reciver function
	fmt.Println("Area of shape :", s.area())
	fmt.Println("Perniter of shape :", s.perimeter())
}

func main() {

	cr := circle{20}
	rt := rectangle{10, 20}

	display(cr)
	display(rt)

}
