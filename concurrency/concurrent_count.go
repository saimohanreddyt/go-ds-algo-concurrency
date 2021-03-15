package main

import (
	"fmt"
	"time"
)

func main() {
	c := make(chan string)

	go count("sheep", c)
	//var wg sync.WaitGroup
	//wg.Add(1)
	//go func() {
	//	count("sheep")
	//	wg.Done()
	//}()
	//go count("fish")
	//fmt.Scanln()
	//time.Sleep(time.Second * 2)
	//wg.Wait()
	for msg := range c {

		fmt.Println(msg)
	}
}
func count(thing string, c chan string) {

	for i := 1; i <= 5; i++ {
		c <- thing
		//fmt.Println(i, thing)
		time.Sleep(time.Millisecond * 500)
	}
	close(c)

}
