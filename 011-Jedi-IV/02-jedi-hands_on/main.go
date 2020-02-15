package main

import "fmt"

func main() {
	x := []int{11, 88, 13, 22, 17, 7, 12, 29, 50, 32}
	for i, v := range x {
		fmt.Println(i, v)
	}
	fmt.Printf("%T\n", x)
}
