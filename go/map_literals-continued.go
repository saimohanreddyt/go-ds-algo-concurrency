package main

import "fmt"

type vertex struct {
	lat, long float64
}

var m = map[string]vertex{
	//If the top-level type is just a type name, you can omit it from the elements of the literal.
	"Bell Labs": {40.68433, -74.39967},
	"Google":    {37.42202, -122.08408},
}

func main() {
	fmt.Println(m)
}
