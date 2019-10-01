/*
A slice has both a length and a capacity

The length of a slice is the number of elements it contains


The capacity of a slice is the number of elements in the underlying array, counting from the first element in the slice.
*/

package main

import (
	"fmt"
)

func main() {
	s := []int{2, 3, 5, 7, 11, 13}
	printSlice(s)

	// slicing the slice to give it a zero length
	s = s[:0]
	printSlice(s)

	// Extending slice length
	s = s[:4]
	printSlice(s)

	//Drop its first two values
	s = s[2:]
	printSlice(s)
}

func printSlice(s []int) {
	fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)
}
