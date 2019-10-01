package main

import (
	"fmt"
)

// func main() {
// 	sum := 0
// 	for i := 9; i < 10; i++ {
// 		sum += i
// 	}
// 	fmt.Println(sum)
// }

// for-continued.go

// func main() {
// 	sum := 1000
// 	for ;sum < 1000; {
// 		sum += sum
// 	}
// 	fmt.Println(sum)
// }

// for-is-gos-while.go
func main() {
	sum := 1
	for sum < 1000 {
		sum += sum
	}
	fmt.Println(sum)
}
