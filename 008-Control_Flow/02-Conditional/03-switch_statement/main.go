package main

import "fmt"

func main() {
	switch {
	case false:
		fmt.Println("this should not print")
	case (2 == 4):
		fmt.Println("this should not printer either")
	case (3 == 3):
		fmt.Println("prints")
		fallthrough
	case (4 == 4):
		fmt.Println("also true, does it print?")
	case (7 == 9):
		fmt.Println("not true")
		fallthrough
	case (11 == 14):
		fmt.Println("not true either")
		fallthrough
	case (15 == 15):
		fmt.Println("true 15")
	}
}
