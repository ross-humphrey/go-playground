package main

/**
Support for times and durations
*/

import (
	"fmt"
	"time"
)

func main() {
	p := fmt.Println

	now := time.Now() // time now
	p(now)

	// can build a time strut
	then := time.Date(2009, 11, 17, 20, 34, 58, 651387237, time.UTC)
	p(then)

	// extract various components of the time value as expected
	p(then.Year())
	p(then.Month())
	p(then.Day())
	p(then.Hour())
	p(then.Minute())
	p(then.Second())
	p(then.Nanosecond())
	p(then.Location())

	p(then.Weekday()) // Monday-Sunday

	//comparing methods
	p(then.Before(now))
	p(then.After(now))
	p(then.Equal(now))

	diff := now.Sub(then) // sub returns a duration representing the interval between two times
	p(diff)

	//compute length of the duration in various units
	p(diff.Hours())
	p(diff.Minutes())
	p(diff.Seconds())
	p(diff.Nanoseconds())

	//Use add tp advange time (-) to move back
	p(then.Add(diff))
	p(then.Add(-diff))
}
