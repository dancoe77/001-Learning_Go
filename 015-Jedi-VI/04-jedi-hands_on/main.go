package main

import "fmt"

type person struct {
	first string
	last  string
	age   int
}

func (p person) speak() {
	fmt.Println(p.first, p.last, p.age)
}

func main() {
	pa1 := person{
		"James",
		"Bond",
		42,
	}
	fmt.Println(pa1)
}
