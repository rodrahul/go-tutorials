package main

import (
	"fmt"
	"io/ioutil"
)

func main() {
	data, err := ioutil.ReadFile("/Users/rrode/go/src/gopdf/readFile.txt")
	if err != nil {
		fmt.Println("Failed to read file with error :", err)
		return
	}
	fmt.Println("File contents: ")
	fmt.Println(string(data))
}