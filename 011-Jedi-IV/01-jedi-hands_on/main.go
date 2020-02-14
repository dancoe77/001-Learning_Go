package main

import "fmt"

func main() {
	x := [5]int{11, 88, 13, 22, 17}
	for i, v := range x {
		fmt.Println(i, v)
	}
	fmt.Printf("%T\n", x)
}
