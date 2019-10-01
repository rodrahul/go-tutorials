package main

import (
	"fmt"
	"math/rand"
	"sort"
)

func main() {
	scores := make([]int, 100)
	fmt.Print("Len and cap of scores slice is:", len(scores), cap(scores), "\n")

	for i := 0; i < 100; i++ {
		scores[i] = int(rand.Int31n(1000))
	}
	sort.Ints(scores)

	worst := make([]int, 5)
	copy(worst, scores[:5])

	fmt.Println(worst)
	fmt.Println(scores)
}
