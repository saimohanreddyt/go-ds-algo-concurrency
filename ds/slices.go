package main

import "fmt"

func main() {
	p := [6]int{1, 2, 3, 4, 5, 6}
	//An array has a fixed size.
	//A slice, on the other hand, is a dynamically-sized, flexible view into the elements of an array.
	//In practice, slices are much more common than arrays.
	var s []int = p[1:4]
	//The type []T is a slice with elements of type T.
	fmt.Println(s)
}
