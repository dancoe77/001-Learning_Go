package main

import "fmt"

func main() {
	foo()
}

//func (r receiver) identifier(parameters) (return(s)) { ... }

func foo() {
	fmt.Println("Hello from Foo")
}
