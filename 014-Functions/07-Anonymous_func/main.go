package main

import "fmt"

func main() {
	foo()

	func() {
		fmt.Println("Anonymous func ran")
	}()
}

func foo() {
	fmt.Println("foo ran")
}
