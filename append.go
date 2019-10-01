package main

import "fmt"

func main() {
	var s []int

	// Append works on nil slices
	s = append(s, 0)
	// fmt.Printf("%p\n", s)
	printSlice(s)

	//The slice grows as needed
	s = append(s, 1)
	printSlice(s)

	s = append(s, 2, 3, 4)
	printSlice(s)
}

func printSlice(s []int) {
	fmt.Printf("%p\n", s)
	fmt.Printf("len=%d cap=%d %v\n\n", len(s), cap(s), s)
}
