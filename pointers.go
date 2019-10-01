package main

import (
	"fmt"
)

func main() {
	i, j := 42, 2701

	p := &i // Point to i
	fmt.Println(*p)
	fmt.Printf("Type of p %T\n", p)
	fmt.Printf("Address of i i.e. &i %p\n", &i)
	fmt.Printf("Value inside p %x\n", p)
	fmt.Printf("Address of p i.e. &p %p\n", &p)

	*p = 21
	fmt.Println(i)

	p = &j
	*p = *p / 37
	fmt.Println(j)
}
