/*
Function with a pointer argument must take a pointer

func ScaleFunc(v *Vertex, f float64) {}

var v Vertex
ScaleFunc(v, 5) 	// Compile error
ScaleFunc(&v, 5)	// OK

While methods with pointer recievers take either a value or a pointer as the reciever when they are called

func (v *Vertex) Scale(f float64) {}

var v Vertex
v.Scale(5)		//OK
p :=&v
p.Scale(5)		//OK

For v.Scale(5), even though v is not a pointer, the method with the pointer reciever is called automaticlly.
GO interprets the statement v.Scale(5) as (&v).Scale(5) since the Scale method has a pointer reciever
*/

package main

import "fmt"

type Vertex struct {
	X, Y float64
}

func ScaleFunc(v *Vertex, f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}

func (v *Vertex) Scale(f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
	fmt.Println("In Scale fn, vertex is, ", v)
}

func main() {
	v := Vertex{3, 4}
	p := Vertex{3, 4}

	ScaleFunc(&v, 10)
	p.Scale(10)

	fmt.Println(v, p)
}
