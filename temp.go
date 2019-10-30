package main

import "fmt"

type A struct {
	a int
}

type B struct {
	b int
}

type myInt1 int
type myInt2 int

func main() {
	a := A{1}
	b := B{2}
	m1 := myInt1(1)
	m2 := myInt2(2)

	fmt.Printf("Type of A:%T Value: %+v\n", a, a)
	fmt.Printf("Type of A:%T Value: %+v\n", b, b)

	fmt.Printf("Type of m1:%T Value: %+v\n", m1, m1)
	fmt.Printf("Type of m2:%T Value: %+v\n", m2, m2)

	m1 = myInt1(m2)
	fmt.Printf("Type of m1:%T Value: %+v\n", m1, m1)

}
