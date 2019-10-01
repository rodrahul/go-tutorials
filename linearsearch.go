package main

import (
	"fmt"
)

func linearSearch(sliceToSearch []int, numToSearch int) bool {
	for _, value := range sliceToSearch {
		if value == numToSearch {
			return true
		}
	}
	return false
}

func main() {
	items := []int{95, 78, 46, 58, 45, 86, 99, 251, 320}
	fmt.Println(linearSearch(items, 45))

}
