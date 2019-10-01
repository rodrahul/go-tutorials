package main

import (
	"fmt"
)

func binarySearch(sliceToSearch []int, numToSearch int) bool {

	left := 0
	right := len(sliceToSearch) - 1

	for left <= right {
		midPoint := (left + right) / 2
		if sliceToSearch[midPoint] < numToSearch {
			left = midPoint + 1
		} else {
			right = midPoint - 1
		}
	}
	if left == len(sliceToSearch) || sliceToSearch[left] != numToSearch {
		return false
	}
	return true

}

func main() {
	items := []int{1, 2, 9, 20, 31, 45, 63, 70, 100}
	fmt.Println(binarySearch(items, 21))

}
