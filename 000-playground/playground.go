package main

import (
	"fmt"
)

func main() {
n, err := fmt.Println("Hello, playground", 42, true)
fmt.Println(n)
fmt.Println(err)
}