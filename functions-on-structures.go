// We can associate a method with a structure

package main

import (
	"fmt"
)

type saiyan struct {
	Name  string
	Power int
}

// *saiyan is the receiver of the Super method
func (s *saiyan) super() {
	s.Power += 10000
}

func main() {
	goku := &saiyan{"Goku", 9001}
	rahul := saiyan{"Rahul", 32}
	rahul.super()
	goku.super()
	fmt.Println(goku.Power)
	fmt.Println(rahul.Power)
}
