package main

import "fmt"

func main() {
	m := make(map[string]int)
	////Insert or update an element in map m: m[key] = elem
	m["Answer"] = 42
	fmt.Println("The value:", m["Answer"])

	//Retrieve an element: elem = m[key]
	m["Answer"] = 48
	fmt.Println("The value:", m["Answer"])

	//Delete an element: delete(m, key)
	delete(m, "Answer")
	fmt.Println("The value:", m["Answer"])
	//Test that a key is present with a two-value assignment: elem, ok = m[key]
	v, ok := m["Answer"]
	//If key is in m, ok is true. If not, ok is false.
	fmt.Println("The value:", v, "Present?", ok)
} //If key is not in the map, then elem is the zero value for the map's element type.
