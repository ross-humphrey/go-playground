package main

// Errors in GO are handled by a separate explicit return value.

import "fmt"
import "errors"

func f1(arg int) (int, error) {
	if arg == 42 {
		return -1, errors.New("can't work with 42") // Construct a basic error value with the error message
	}

	return arg + 3, nil // nil indicates no error
}

type argError struct {
	arg  int
	prob string
}

func (e *argError) Error() string { // Custom errors can be implemented by implementing the error method on them
	return fmt.Sprintf("%d - %s", e.arg, e.prob)
}

func f2(arg int) (int, error) {
	if arg == 42 {
		return -1, &argError{arg, "can't work with it"} // use &argError to build a new struct supplying values
	}
	return arg + 3, nil
}

func main() {

	for _, i := range []int{7, 42} {
		if r, e := f1(i); e != nil {
			fmt.Println("f1 failed", e)
		} else {
			fmt.Println("f1 worked", r)
		}
	}

	for _, i := range []int{7, 42} {
		if r, e := f2(i); e != nil {
			fmt.Println("f2 failed", e)
		} else {
			fmt.Println("f2 worked", r)
		}
	}

	_, e := f2(42) //programmatically use data in custom error - get error as instance of custom error type via assertion
	if ae, ok := e.(*argError); ok {
		fmt.Println(ae.arg)
		fmt.Println(ae.prob)
	}
}
