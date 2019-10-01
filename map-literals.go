// Map literals are like struct literals, but the keys are required.

package main

import (
	"fmt"
)

type Vertex struct {
	Lat, Long float64
}

var m = map[string]Vertex{
	"Bell Labs": Vertex{
		40.68433, -74.39967,
	},
	"Google": Vertex{
		37.42202, -122.08408,
	},
}

var m2 = map[string]Vertex{
	"Bell Labs": {40.68433, -74.39967},
	"Google":    {37.42202, -122.08408},
}
var m3 = map[int]string{
	1: "one",
	2: "two",
}

func main() {
	fmt.Println(m)
	fmt.Println(m2)
	fmt.Println(m3)
}
