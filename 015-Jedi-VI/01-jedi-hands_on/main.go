package main

import "fmt"

func main() {
	foo(35)
	bar("James Bond")
}

func foo(i int) {
	fmt.Println("Hello from foo, ", i)
}

func bar(st string) (x int) {
	a := fmt.Sprint(st, "is")
	b := 42
	fmt.Println(a, b)
}
