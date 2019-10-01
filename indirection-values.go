/*
Functions that take value argument must take a value of that specific type

func AbsFunc(v Vertex) float64{}

var v Vertext
AbsFunc(v)			// OK
AbsFunc(&v)			// Compile Error


While methods with value recievers take either a value or a pointer as the reciever when they are called

func (v Vertext) Abs() float64 {}

var v Vertext
v.Abs()		//Ok
p := &v
p.Abs()		//OK
In this case, the method call p.Abs() is interpreted as (*p).Abs().
*/
package main

import (
	"fmt"
	"math"
)

type Vertex struct {
	X, Y float64
}

func (v Vertex) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func AbsFunc(v Vertex) float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func (v Vertex) Scale(f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
	fmt.Println("In Scale fn, vertex is, ", v)
}

func main() {
	v := Vertex{3, 4}
	fmt.Println(v.Abs())
	fmt.Println(AbsFunc(v))
	v.Scale(10)
	fmt.Println(v)

	p := &Vertex{4, 3}
	fmt.Println(p.Abs())
	fmt.Println(AbsFunc(*p))
	p.Scale(10)
	fmt.Println(p)

}
