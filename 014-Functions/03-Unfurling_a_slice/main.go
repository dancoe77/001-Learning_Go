package main

import "fmt"

func main() {
	x := sum("james")
	fmt.Println("The total is", x)
}

func sum(s string, x ...int) int {
	fmt.Println(x)
	fmt.Printf("%T\n", x)

	sum := 0
	// var sum int
	for i, v := range x {
		sum += v
		fmt.Println("for item in index position", i, "we are now adding", v, "to the total which is now", sum)
	}
	fmt.Println("The total is", sum)
	return sum
}

// func (r receiver) identifier(parameter(s)) (return(s)) { code}
