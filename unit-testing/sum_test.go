package main

import "testing"

// Learning from: https://blog.alexellis.io/golang-writing-unit-tests/

func TestSum(t *testing.T) { // first and only param is *testing. Starts with Test followed by Caps
	total := Sum(5, 5)
	if total != 10 {
		t.Errorf("Sum was incorrect, got: %d, want: %d", total, 10) // calls t.Error or t.Fail
	}
	// t.log can be used to provide non-failing debug info
}

/**
Why there are no assertions in Go

> tests can feel like they are written in a different language
> errors can be cryptic
> pages of stack traces
> tests stop after first assert fails - masking patterns of failure

*/
