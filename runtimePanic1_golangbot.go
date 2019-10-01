package main

import (
	"fmt"
)

/*
recover is a builtin fun which is used to regain control of a panicking goroutine
func recover() interface{}

*/
func recoverFromA() {
	if r:= recover(); r!=nil {
		fmt.Println("Recovered from func a: ", r)
		//debug.PrintStack()
	}
}

func a() {
	defer recoverFromA()
	n := []int{1,2,3}
	fmt.Println(n[3])
	fmt.Println("Normally returned from a")
}

func main() {
	a()
	fmt.Println("Normally returned from main")
}
