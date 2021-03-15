package main

import (
	"fmt"
)

func greatsmall(x, y int) int {

	//Go's if statements are like its for loops; the expression need not be surrounded by parentheses ( ) but the braces { } are required.

	if x > y {
		fmt.Println("x is greater")

		return x
	}
	if x < y {
		fmt.Println("Y is greater")
		return y
	}

	return greatsmall(x, y)
}

func main() {
	greatsmall(10, 20)
}
