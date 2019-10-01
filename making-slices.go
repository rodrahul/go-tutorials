/*
Slices can be created with the built-in make function; this is how you create dynamically-sized arrays.

The make function allocates a zeroed array and returns a slice that refers to that array
*/

package main

import "fmt"

func main() {
	a := make([]int, 5)
	printSlice("a", a)

	b := make([]int, 0, 5)
	printSlice("b", b)

	c := b[:2]
	printSlice("c", c)

	d := c[2:5]

	printSlice("d", d)
	fmt.Println(a)
	fmt.Println(b)

	fmt.Printf("Address of b is %p\n", b)
	fmt.Printf("Address of c is %p\n", c)
	fmt.Printf("Address of d is %p\n", d)
	fmt.Printf("Address of d is %p\n", &a[0])
	fmt.Printf("Address of d is %p\n", &a[1])
	fmt.Printf("Address of d is %p\n", &a[2])
}

func printSlice(s string, x []int) {
	fmt.Printf("%s len=%d cap=%d %v\n",
		s, len(x), cap(x), x)
}
