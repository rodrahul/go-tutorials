package main

import (
	"fmt"
)

func main() {
	// var country map[int]string
	country := make(map[int]string)

	country[0] = "India"
	country[1] = "China"
	country[2] = "USA"
	country[3] = "Germany"
	fmt.Println(country)

	for key, value := range country {
		fmt.Printf("Key: %d Value: %s\n", key, value)
	}
}
