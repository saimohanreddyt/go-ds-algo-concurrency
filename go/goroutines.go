package main

import (
	"fmt"
	"time"
)

func helo(s string) {
	time.Sleep(10 * time.Millisecond)
	fmt.Println(s)
}
func main() {
	go helo("World")

	helo("Hello")
}
