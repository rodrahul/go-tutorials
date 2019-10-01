package main

import (
	"fmt"
	"math"
)

func pow(x, n, lim float64) float64 {
	if v := math.Pow(x, n); v < lim {
		return v
	}
	return x
}

func main() {
	fmt.Printf("Type of %T \n", pow(3, 3, 10))

}
