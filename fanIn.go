/*
Taking values from many channels, and putting those values onto one channel.
*/
package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	c := fanIn(boring("Rahul"), boring("Supriya"))
	fmt.Println("after calling fanIN")

	for i := 0; i < 20; i++ {
		fmt.Println(<-c)
	}
	fmt.Println("You're both boring; I'm leaving.")
}

func boring(msg string) <-chan string {
	c := make(chan string)

	go func() {
		for i := 0; ; i++ {
			c <- fmt.Sprintf("%s %d", msg, i)
			time.Sleep(time.Duration(rand.Intn(1e3)) * time.Millisecond)
		}
	}()
	fmt.Println("Returning from boring fn ", msg)
	return c
}

// Takes 2 args, receiving channels
// returns receiving channel
func fanIn(input1, input2 <-chan string) <-chan string {
	c := make(chan string)

	go func() {
		for {
			c <- <-input1 //receive from input1 and send to c
		}
	}()

	go func() {
		for {
			c <- <-input2
		}
	}()

	return c

}
