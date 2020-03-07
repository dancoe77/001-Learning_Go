package main

import "fmt"

func main() {
	x := bar()

	fmt.Printf("%T\n", x)

	fmt.Println(x())
}

func foo() string {
	return "Hello World"
}

func bar() func() int {
	return func() int {
		return 451
	}
}
