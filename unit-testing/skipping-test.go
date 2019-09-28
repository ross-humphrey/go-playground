package main

import "testing"

// Tests or benchmarks may be skipped at run time with a call to the Skip method of *T or *B

func TestTimeConsuming(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping test in short mode.")
	}
}
