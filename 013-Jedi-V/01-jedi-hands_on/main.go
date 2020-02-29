package main

import "fmt"

type person struct {
	first      string
	last       string
	favFlavors []string
}

func main() {
	p1 := person{
		first: "James",
		last:  "Bond",
		favFlavors: []string{
			"Chocolate",
			"Martini",
			"Rum and Coke",
		},
	}

	p2 := person{
		first: "Miss",
		last:  "Moneypenny",
		favFlavors: []string{
			"Vanilla",
			"Strawberry",
			"Capuccino",
		},
	}

	fmt.Println(p1.first)
	fmt.Println(p1.last)
	for i, v := range p1.favFlavors {
		fmt.Println(i, v)
	}
	fmt.Println(p2.first)
	fmt.Println(p2.last)
	for j, w := range p2.favFlavors {
		fmt.Println(j, w)
	}
}
