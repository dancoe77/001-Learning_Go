package main

import "fmt"

func main() {
	x := []int{42, 43, 44, 45, 46, 47, 48, 49, 50, 51}
	fmt.Println(x[0:4])
	fmt.Println(x[5:9])
	fmt.Println(x[2:6])
	fmt.Println(x[1:5])
}
