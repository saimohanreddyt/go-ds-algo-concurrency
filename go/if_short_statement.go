package main

import (
	"fmt"
	"math"
)

func pow(x, n, l float64) float64 {
	//Like for, the if statement can start with a short statement to execute before the condition.
	if v := math.Pow(x, n); v < l {
		return v
	}
	return l
}

//Variables declared by the statement are only in scope until the end of the if.
func main() {
	fmt.Println(
		pow(3, 2, 10),
		pow(3, 3, 20),
	)
}
