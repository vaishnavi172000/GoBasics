package main

import (
	"errors"
	"fmt"
)

// A sentinel error is a predeclared variable that is used to signify a specific error condition.
var ErrorOddNumber = fmt.Errorf("given number is odd")

//By convention, errors are the last return value and have type error, a built-in interface.

// Go’s approach makes it easy to see which functions return errors
// and to handle them using the same language constructs employed for other, non-error tasks.
func f(i int) (int, error) {
	if i%2 == 0 {
		return int(i / 2), nil
	}
	return 0, ErrorOddNumber

}

// errors.New constructs a basic error value with the given error message.
func testError(i int) (int, error) {
	if i == 1 {
		return i, nil
	} else {
		if i%2 == 0 {
			return 0, errors.New("input value is 2")
		} else {
			//We can wrap errors with higher-level errors to add context.
			// The simplest way to do this is with the %w verb in fmt.Errorf.
			//  Wrapped errors create a logical chain (A wraps B, which wraps C, etc.)
			//  that can be queried with functions like errors.Is and errors.As.
			return 0, fmt.Errorf("input is invalid: %w", ErrorOddNumber)
		}
	}
}

func main() {
	for i := 1; i <= 2; i++ {
		fmt.Println(f(i))
	}
	for i := 1; i <= 3; i++ {
		if v, err := testError(i); err != nil {
			// errors.Is checks that a given error (or any error in its chain)
			// matches a specific error value.
			// This is especially useful with wrapped or nested errors,
			//  allowing you to identify specific error types or sentinel errors in a chain of errors.
			if errors.Is(err, ErrorOddNumber) {
				fmt.Println("this error is wrapped eror of odd number err", err)
			} else {
				fmt.Println("Error from function", err)
			}
		} else {
			fmt.Println("values returned :", v)
		}
	}

}
