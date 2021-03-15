/*
Goroutine: A Goroutine is a function or method which executes independently and simultaneously in connection with any other Goroutines present in your program.
Or in other words, every concurrently executing activity in Go language is known as a Goroutines.
*/
package main

import (
	"fmt"
	"time"
)

func hello() {
	fmt.Println("Hello world goroutine")
}
func main() {
	go hello()
	time.Sleep(3 * time.Second)
	fmt.Println("main function")
}
