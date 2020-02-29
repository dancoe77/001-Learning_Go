package main

import "fmt"

func main() {
	tr := struct {
		doors int
		color string
	}{
		doors: 2,
		color: "blue",
	}
	fmt.Println(tr)
}
