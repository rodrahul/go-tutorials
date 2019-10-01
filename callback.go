/*
passing a func as an argument
functional programming not something that is recommended in go, however, it is good to be aware of callbacks
idiomatic go: write clear, simple, readable code

*/
package main

import (
	"fmt"
)

func sum(x ...int) int {
	var total int
	for _, v := range x {
		total += v
	}
	fmt.Printf("In sum type of argument is %T\n", x)
	return total
}

func even(f func(xi ...int) int, vi ...int) int {
	var yi []int
	for _, v := range vi {
		if v%2 == 0 {
			yi = append(yi, v)
		}
	}
	return f(yi...)
}

func main() {
	ii := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	s := sum(ii...)
	fmt.Println("all numbers", s)

	s = even(sum, ii...)
	fmt.Println("even numbers", s)
}
