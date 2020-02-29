package main

import "fmt"

type vehicle struct {
	doors string
	color string
}

type truck struct {
	vehicle
	fourWheel bool
}

type sedan struct {
	vehicle
	luxury bool
}

func main() {
	tr := truck{
		vehicle: vehicle{
			doors: "Extended",
			color: "Brown",
		},
		fourWheel: true,
	}

	sd := sedan{
		vehicle: vehicle{
			doors: "Regular",
			color: "Blue",
		},
		luxury: true,
	}

	fmt.Println(tr)
	fmt.Println(sd)

	fmt.Println(tr.doors)
	fmt.Println(sd.color)
}
