package main

import "fmt"

func main() {
	x := 33
	for {
		if x > 123 {
			break
		}
		fmt.Printf("%d\t%u\n", x, x)
		x++
	}
	fmt.Println("done.")
}
