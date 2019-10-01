/*
Method is a function with a special reciever argument
*/

package main

import (
	"fmt"
	"math"
)

type Vertex struct {
	X, Y float64
}

// Abs method has a reciever of type Vertex named v
// A method is just a function with a reciever argument
func (v Vertex) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

// Abs as a regular function
func AbsFunc(v Vertex) float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func main() {
	v := Vertex{3, 4}
	fmt.Println(v.Abs())
	fmt.Println(AbsFunc(v))
}
