package main

import "fmt"

//func swap(a string , b string) (string, string)  instead
// we can use whentwo or more consecutive named function parameters share a type, you can omit the type from all but the last.
func swap(a, b string) (string, string) {

	return b, a
}

func main() {
	a, b := swap("world", "hello")

	fmt.Println(a, b)

}
