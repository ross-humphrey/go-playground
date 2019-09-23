package main

// How to get time since Unix epoch

import (
	"fmt"
	"time"
)

func main() {

	// use time.Now with Unix or UnixNano to get elapsed time since the Unix epoch
	now := time.Now()
	secs := now.Unix()
	nanos := now.UnixNano()
	fmt.Println(now)

	// no UnixMillis -
	millis := nanos / 10000000
	fmt.Println(secs)
	fmt.Println(millis)
	fmt.Println(nanos)

	// convert integer seconds or nanoseconds since epoch
	fmt.Println(time.Unix(secs, 0))
	fmt.Println(time.Unix(0, nanos))
}
