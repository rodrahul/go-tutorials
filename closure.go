package main

import (
	"fmt"
)

// package level scope

func main() {
	a := incrementor()
	b := incrementor()

	fmt.Printf("Type is %T, value is %v", a, a)

	fmt.Println(a())
	fmt.Println(a())
	fmt.Println(a())
	fmt.Println(a())
	fmt.Println(b())
	fmt.Println(b())
}

func incrementor() func() int {
	var x int
	fmt.Printf("%p\n", &x)
	return func() int {
		x++
		return x
	}
}

/*
Closure is when we have “enclosed” the scope of a variable in some code block.
*/
