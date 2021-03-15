package main

import (
	"fmt"
	"math/cmplx"
)

var (
	T bool       = false
	M uint64     = 1<<64 - 1
	z complex128 = cmplx.Sqrt(-5 + 12i)
)

func main() {
	fmt.Printf("Type: %T Value: %v\n", T, T)
	fmt.Printf("Type: %T Value: %v\n", M, M)
	fmt.Printf("Type: %T Value: %v\n", z, z)
}
