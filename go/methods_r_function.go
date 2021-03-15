package main

import (
	"fmt"
	"math"
)

//Remember: a method is just a function with a receiver argument.
type Vertex struct {
	X, Y float64
}

//Here's Abs written as a regular function with no change in functionality.
func Abs(v Vertex) float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func main() {
	v := Vertex{3, 4}
	fmt.Println(Abs(v))
}
