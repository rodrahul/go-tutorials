/*
A type switch is a construct that permits serveral type assertions in series

A type switch is like a regular switch statement, but the cases in a type switch specify type (not values), and those values are compared against the type of the value held by the given interface value
*/

package main

import (
	"fmt"
)

func do(i interface{}) {
	switch v := i.(type) {
	case int:
		fmt.Printf("Twice of %v is %v\n", v, v*2)
	case string:
		fmt.Printf("%q is %v bytes long\n", v, len(v))
	default:
		fmt.Printf("I don't know about type %T!\n", v)
	}
}

func main() {
	do(21)
	do("hello")
	do(true)
}
