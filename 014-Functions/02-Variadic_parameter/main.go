package main

import "fmt"

func main() {
	foo()
}

func foo(x ...int) {
	fmt.Println(x)
	fmt.Printf("%T\n", x)
}

// func (r receiver) identifier(parameter(s)) (return(s)) { code}
