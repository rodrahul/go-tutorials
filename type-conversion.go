// The expression T(v) converts v to the type T

// GO assignment between items of different type requires an explicit conversion

package main

import (
	"fmt"
	"math"
)

func main() {
	v := "Rahul" //Change me
	var x, y int = 3, 4
	var f float64 = math.Sqrt(float64(x*x + y*y))
	var z uint = uint(f)
	fmt.Println(x, y, z)
	fmt.Printf("v is of type %t\n", v)
}
