package main

import "fmt"

func main() {
	a := 42
	fmt.Println(a)
	fmt.Println(&a) // & gives you the address

	fmt.Printf("%T\n", a)
	fmt.Printf("%T\n", &a)

	//var b int = &a {can't use *int as int}
	var b *int = &a
	fmt.Println(b)
	fmt.Println(*b) // * gives me the value stored at an address
	fmt.Println(*&a)
	//b := &a

	*b = 43
	fmt.Println(a)
}
