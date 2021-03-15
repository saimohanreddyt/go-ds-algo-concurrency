package main

import (
	"fmt"
	"time"
)

func main() {
	c1 := make(chan string)
	c2 := make(chan string)
	go func() {
		for {
			c1 <- "every 500ms"
			time.Sleep(time.Millisecond * 500)
		}
	}()

	go func() {
		for {
			c2 <- "Every two seconds"
			time.Sleep(time.Second * 2)
		}
	}()
	for i := 1; i <= 5; i++ {
		fmt.Println(<-c1)
		fmt.Println(<-c2)
	}
}
