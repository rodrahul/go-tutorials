/*
A nil interface value holds neither value nor concrete type

Calling a method on a nil interface isa a run-time error because there is not type insdie the interface tuple to indicate which concrete method to call
*/

package main

import (
	"fmt"
)

type I interface {
	M()
}

func describe(i I) {
	fmt.Println("%v, %T\n", i, i)
}

func main() {
	var i I
	describe(i)
	i.M()

}
