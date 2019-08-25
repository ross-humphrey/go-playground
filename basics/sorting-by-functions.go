package main

// Sort a collection by something other than natural order
// Example of custom sorts

import "sort"
import "fmt"

// alias for builtin []string type.
type byLength []string

// implement sort.Interface -Len, Less and Swap on type
// to use sort package genric Sort function
func (s byLength) Len() int {
	return len(s)
}

func (s byLength) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func (s byLength) Less(i, j int) bool {
	return len(s[i]) < len(s[j])
}

func main() {
	fruits := []string{"peach", "banana", "kiwi"}
	sort.Sort(byLength(fruits))
	fmt.Println(fruits)
}
