package main

import "fmt"

func main() {
	fmt.Println("Hello, playground")

	f := func() {
		fmt.Println("my first func expression")
	}
	f()
}
