package main

import "fmt"

func main() {

	a := 1024 == 1024
	b := 512 <= 1024
	c := 1024 >= 512
	d := 512 != 1024
	e := 512 < 1024
	f := 1024 > 512
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
	fmt.Println(d)
	fmt.Println(e)
	fmt.Println(f)
}
