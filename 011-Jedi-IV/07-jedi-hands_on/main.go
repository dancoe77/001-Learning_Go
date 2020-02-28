package main

import "fmt"

func main() {
	jb := []string{"James", "Bond", "Shaken, not stirred"}
	fmt.Println(jb)
	mp := []string{"Miss", "Moneypenny", "Helloooooo, James."}
	fmt.Println(mp)

	for i, v := range jb {
		fmt.Println(i, v)
	}
	for i, v := range mp {
		fmt.Println(i, v)
	}
}
