package main

import "fmt"

func main() {
	i, j := 33, 99
	//Go has pointers. A pointer holds the memory address of a value.
	p := &i

	//The type *T is a pointer to a T value. Its zero value is nil.
	fmt.Println(*p)
	*p = 21
	fmt.Println(i)

	//The & operator generates a pointer to its operand.
	p = &j
	//The * operator denotes the pointer's underlying value.
	*p = *p / 37
	fmt.Println(j)
}
