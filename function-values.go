package main

import (
	"fmt"
	"math"
)

func compute(fn func(float64, float64) float64, a, b float64) float64 {
	return fn(a, b)
}

func main() {
	hypot := func(x, y float64) float64 {
		return math.Sqrt(x*x + y*y)
	}

	fmt.Println(hypot(5, 12))

	fmt.Println(compute(hypot, 3, 4))
	fmt.Println(compute(math.Pow, 3, 4))
}
