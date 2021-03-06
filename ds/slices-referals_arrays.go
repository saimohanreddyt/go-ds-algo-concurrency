package main

import "fmt"

func main() {
	names := [4]string{
		"h",
		"a",
		"n",
		"i",
	}
	fmt.Println(names)
	//A slice does not store any data, it just describes a section of an underlying array.

	a := names[0:2]
	b := names[1:3]
	fmt.Println(a, b)
	//Changing the elements of a slice modifies the corresponding elements of its underlying array.
	//Other slices that share the same underlying array will see those changes.
	b[0] = "yes"
	fmt.Println(a, b)
	fmt.Println(names)
}
