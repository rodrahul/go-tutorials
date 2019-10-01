/*
We can also declare methods with pointer receivers.

Methods with pointer receivers can modify the value to which the receiver points.
*/

package main

import (
	"fmt"
	"math"
)

type Vertex struct {
	X, Y float64
}

// scale method is defined on *Vertex, the scale operation will change the values since its working with the pointer and not copy
func (v *Vertex) Scale(f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}

func (v Vertex) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func main() {
	v := Vertex{3, 4}
	v.Scale(10)
	fmt.Println(v.Abs())
}
