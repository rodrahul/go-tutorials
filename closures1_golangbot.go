package main

import (
	"fmt"
)

// appendStr returns a function, which accepts a string and returns a string
func appendStr() func(string) string {
	t := "Hello"
	c := func(b string) string {
		t = t + " " + b
		return t
	}
	return c
}

func main() {
	a := appendStr()
	b := appendStr()
	
	fmt.Println(a("World"))
	fmt.Println(b("Everyone"))

	fmt.Println(a("Gopher"))
	fmt.Println(b("!"))
	
	// For my testing
	fmt.Printf("Type of a is %T\n", a)
}

/*

Closures are a special case of anonymous functions. Closures are anonymous functions which access the variables defined outside the body of the function

*/