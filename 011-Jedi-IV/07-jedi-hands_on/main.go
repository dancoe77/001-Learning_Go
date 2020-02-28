package main

import "fmt"

func main() {
	jb := []string{"James", "Bond", "Shaken, not stirred"}
	fmt.Println(jb)
	mp := []string{"Miss", "Moneypenny", "Helloooooo, James."}
	fmt.Println(mp)

	jbmp := [][]string{jb, mp}
	fmt.Println(jbmp)

	for i, v := range jbmp {
		fmt.Println("record: ", i)
		for j, val := range v {
			fmt.Printf("\t index position: %v \t value: \t %v \n", j, val)
		}
	}
}
