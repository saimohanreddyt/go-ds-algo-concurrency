package main

import "fmt"

type vertex struct {
	x, y int
}

//A struct literal denotes a newly allocated struct value by listing the values of its fields.
var (
	v1 = vertex{1, 2}  // has type Vertex
	v2 = vertex{x: 1}  // Y:0 is implicit
	v3 = vertex{}      // X:0 and Y:0
	v4 = &vertex{1, 2} // has type *Vertex
)

func main() {
	fmt.Println(v1, v2, v3, v4)

}
