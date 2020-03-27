package main

import "fmt"
import "math"

type circle struct {
	radius float64
}

func (c *circle) area() float64 {
	return math.Pi * c.radius * c.radius
}

type shape interface {
	area() float64
}

func info(s shape) {
	fmt.Println("area", s.area())
}

func main() {
	c := circle{
		radius: 5,
	}

	fmt.Printf("%T\n", &c)
	info(&c)
}
