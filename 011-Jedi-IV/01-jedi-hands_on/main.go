package main

import "fmt"

func main() {
	var x [5]int
	x[0] = 11
	x[1] = 88
	x[2] = 13
	x[3] = 22
	x[4] = 17
	fmt.Println(x)
	fmt.Printf("%T\n", x)
}
