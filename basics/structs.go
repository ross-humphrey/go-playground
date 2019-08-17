package main

import "fmt"

type person struct {
	name string
	age  int
}

func main() {

	fmt.Println(person{"Bob", 20})

	fmt.Println(person{name: "Fred"}) // Omitted fields will be zero valued

	fmt.Println(&person{
		name: "Ann",
		age:  40,
	}) // An & prefix yields a pointer to the struct

	s := person{
		name: "Sean",
		age:  50,
	}
	// Access struct fields with dog
	fmt.Println(s.name)

	sp := &s // yields pointer to s, changes to sp also change s
	fmt.Println(sp.age)

	sp.age = 51
	fmt.Println(sp.age)

	fmt.Println(s.age)
}
