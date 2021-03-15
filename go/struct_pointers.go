package main

import "fmt"

type vertex struct {
	x int
	y int
}

func main() {
	v := vertex{2, 3}
	//To access the field X of a struct when we have the struct pointer p we could write (*p).X.
	//However, that notation is cumbersome, so the language permits us instead to write just p.X, without the explicit dereference.
	p := &v
	p.x = 1e9
	fmt.Println(v)
}
