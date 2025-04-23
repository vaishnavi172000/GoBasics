package main

import (
	"fmt"
	"time"
)

var weekday string

// init function is called before main function
// init function is used to initialize variables or perform setup tasks
// init function is called automatically by go runtime before main function
func init() {
	fmt.Println("init function called")
	// init function is called before main function
	weekday = time.Now().Weekday().String()
}

func main() {
	fmt.Printf("Today is %s", weekday)
}

//init function called
//Today is Monday
