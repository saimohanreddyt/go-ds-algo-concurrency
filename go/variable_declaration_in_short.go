package main

import "fmt"

func main() {
	//Inside a function, the := short assignment statement can be used in place of a var declaration with implicit type.
	a := 1
	b, c, d := "s", "a", "i"
	fmt.Println(a, b, c, d)
}
