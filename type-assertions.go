/*
A type assertion provides access to an interface value's underlying concrete value
Interface values can be thought of as a tuple of a value and concrete type
(value, type)
*/

package main

import (
	"fmt"
)

func main() {
	var i interface{} = "hello"

	s := i.(string)
	fmt.Println(s)

	s, ok := i.(string)
	fmt.Println(s, ok)

	f, ok := i.(float64)
	fmt.Println(f, ok)

	f = i.(float64)
	fmt.Println(f)
}
