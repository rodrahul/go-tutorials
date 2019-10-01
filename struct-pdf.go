package main

import "fmt"

type saiyan struct {
	Name  string
	Power int
}

// Passing values as copies
// func main() {
// 	goku := saiyan{Name: "Goku", Power: 9000}
// 	super(goku)
// 	fmt.Println(goku.Power)
// 	fmt.Printf("Address of goku is %p\n", &goku)
// }

// func super(s saiyan) {
// 	s.Power += 10000
// 	fmt.Printf("Address of s is %p\n", &s)
// }

// --------------------------------------------------

// Passing value by pointers
func main() {
	goku := &saiyan{"Goku", 9001}
	goku.super()
	fmt.Println(goku.Power) // will print 19001
	fmt.Printf("Address of goku is %p\n", goku)
	fmt.Println("Goku is, ", goku)
}

func (s *saiyan) super() {
	// s.Power += 10000
	// Above and below are the same
	//(*s).Power += 10000
	fmt.Println("s is, ", s)
	s = &saiyan{"Gohan", 1000}
	fmt.Printf("Address of s is %p\n", s)
}
