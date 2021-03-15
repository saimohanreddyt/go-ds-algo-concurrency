package main

import "fmt"

func main() {
	message := make(chan string)

	go func() { message <- "IAS" }()

	msg := <-message
	fmt.Println(msg)
}
