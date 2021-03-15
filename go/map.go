package main

import (
	"fmt"
)

type vertex struct {
	lat, long float64
}

//A map maps keys to values.
var m map[string]vertex

//The zero value of a map is nil. A nil map has no keys, nor can keys be added.
func main() {
	//The make function returns a map of the given type, initialized and ready for use.
	m = make(map[string]vertex)
	m["Bell Bottom"] = vertex{
		40.68433, -74.39967,
	}
	fmt.Println(m["Bell Bottom"])
}
