package main

import "fmt"

type vertex struct {
	x int
	y int
}

func main() {
	v := vertex{1, 2}
	fmt.Println(v)

	//Struct fields are accessed using a dot.
	v.x = 4
	fmt.Println(v.x)
	fmt.Println(v)
}
