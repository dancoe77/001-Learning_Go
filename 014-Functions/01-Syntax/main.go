package main

import "fmt"

func main() {
	foo()
	bar("James")
}

//func (r receiver) identifier(parameters) (return(s)) { ... }

func foo() {
	fmt.Println("Hello from Foo")
}

func bar(s string) {
	fmt.Println("Hello,", s)
}
