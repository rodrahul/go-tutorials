package main

import (
	"flag"
	"fmt"
	"os"
)

var (
	flagOne = flag.Int("a", 99, "This is the first flag")
	flagTwo = flag.String("b", "two", "This is the second flag")
	flagThree = flag.Bool("c", false, "This is the second flag")
)


func init() {
	for _, value := range os.Args[1:] {
		fmt.Println(value)
	}
	flag.PrintDefaults()
	flag.Parse()
}


func main() {
	fmt.Println("firstArg", *flagOne)
	fmt.Println("secondArg", *flagTwo)
	fmt.Println("thirdArg", *flagThree)
}