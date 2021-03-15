package main

import "fmt"

//A struct is a collection of fields.
type vertex struct {
	x int
	y int
}

func main() {
	fmt.Println(vertex{1, 3})
}
