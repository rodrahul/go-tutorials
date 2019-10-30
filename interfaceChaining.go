package main

import "fmt"

type Speaker interface {
	Speak() string
}

// Dog implements the Speaker interface
type Dog struct {}

func (d Dog) Speak () string {
	return "Woof woof !!!"
}

// In order to chain this, we need to create a struct that will wrap an existing Speaker and create a chain
type prefixSpeaker struct {
	s Speaker
}

func (ps prefixSpeaker) Speak() string {
	return "prefix: " + ps.s.Speak()
}

func main() {
	ps := prefixSpeaker{s: Dog{}}
	fmt.Println(ps.Speak())
}