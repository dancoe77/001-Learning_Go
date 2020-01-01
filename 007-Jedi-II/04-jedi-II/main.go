package main

import "fmt"

const (
	_  = iota
	kb = 1 << (iota * 10)
	mb = 1 << (iota * 10)
)

func main() {
	fmt.Printf("%d\t\t%b\t\t%#x\n", kb, kb, kb)
	fmt.Printf("%d\t\t%b\t%#x\n", mb, mb, mb)
}
