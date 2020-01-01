package main

import "fmt"

const (
	a = (iota + 2019)
	w = (iota + a)
	x = (iota + a)
	y = (iota + a)
	z = (iota + a)
)

func main() {
	fmt.Printf("%d\t%d\t%d\t%d\n", w, x, y, z)
}
