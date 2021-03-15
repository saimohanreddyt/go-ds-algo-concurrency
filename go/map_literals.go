package main

import "fmt"

type vertex struct {
	lat, long float64
}

//Map literals are like struct literals, but the keys are required.
var m = map[string]vertex{
	"Bell BOttom": vertex{
		40.68433, -74.39967,
	},
	"Google": vertex{
		37.42202, -122.08408,
	},
}

func main() {
	fmt.Println(m)
}
