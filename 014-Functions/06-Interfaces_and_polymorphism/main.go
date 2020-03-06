package main

import "fmt"

type person struct {
	first string
	last  string
}

type secretAgent struct {
	person
	ltk bool
}

// func (r receiver) identifier(parameters) (return(s)) { code}

func (s secretAgent) speak() {
	fmt.Println("I am", s.first, s.last)
}

// var x int
// type foo int
// a VALUE can be of more than one TYPE

type human interface {
	speak()
}

func bar(h human) {
	fmt.Println("I called human")
}

func main() {
	//x := 42
	sa1 := secretAgent{
		person: person{
			"James",
			"Bond",
		},
		ltk: true,
	}

	sa2 := secretAgent{
		person: person{
			"Miss",
			"Moneypenny",
		},
		ltk: true,
	}

	p1 := person{
		first: "Dr.",
		last:  "Yes",
	}

	fmt.Println(sa1)
	sa1.speak()
	sa2.speak()

	fmt.Println(p1)
}
