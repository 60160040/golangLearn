package main

import (
	"fmt"
	"math"
)

func main() {
	c := Circle{0, 0, 5}
	fmt.Println(circleArea(&c))
}

 Circle is circle
type Circle struct {
	x, y, r float64
}

func circleArea(c *Circle) float64 {
	return math.Pi * c.r * c.r
}
