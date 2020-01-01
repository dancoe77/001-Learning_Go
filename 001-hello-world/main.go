package main

import "fmt"

func main() {
	n1, err1 := fmt.Println("Hello everyone, this is the most awesome class for learning the GO programming language", 8, true)
	fmt.Println(n1)
	fmt.Println(err1)
	foo()
	fmt.Println("something else")

	for i := 0; i < 100; i++ {
		if i %2 == 0{
			fmt.Println(i)
		}

	}

	bar()
}

func foo() {
 	n2, err2 := fmt.Println("I'm in foo", 7, false)
 	fmt.Println(n2)
 	fmt.Println(err2)
}

func bar() {
	n3, err3 := fmt.Println("and then we exited", 5, true)
	fmt.Println(n3)
	fmt.Println(err3)
}

// control flow
// (1) sequnece
// (2) loop; iterative
// (3) conditional