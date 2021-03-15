package main

import "fmt"

func main() {
	//array
	a := [10]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	//printing actual array
	fmt.Println(a)
	// slicing the array
	var b []int = a[1:4]
	var c []int = a[0:]
	var d []int = a[:10]
	var e []int = a[:]
	//printing slices arrays
	fmt.Println(b)
	fmt.Println(c)
	fmt.Println(d)
	fmt.Println(e)
}
