package main

import "fmt"

const (
	_  = iota
	kb = 1 << (iota * 10)
)

func main() {
	fmt.Printf("%d\t%b\t%#X\n", kb, kb, kb)
}
