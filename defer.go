package main

import (
	"fmt"
)

// defer.go
// func main() {
// 	defer fmt.Println("world")

// 	fmt.Println("hello")
// }

// defer-multi.go

func main() {
	fmt.Println("counting")

	for i := 0; i < 10; i++ {
		defer fmt.Println(i)
	}

	fmt.Println("done")
}
