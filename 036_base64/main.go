package main

import (
	"encoding/base64"
	"fmt"
	"log"
)

func main() {
	s := "This is the phrase i want to encode to base64"

	s64 := base64.StdEncoding.EncodeToString([]byte(s))

	fmt.Println(len(s))
	fmt.Println(len(s64))
	fmt.Println("String is: ", s)
	fmt.Println("Encoded string is: ", s64)

	// Now Decoding the string
	d64, err := base64.StdEncoding.DecodeString(s64)
	if err != nil {
		log.Fatal("Cannot decode string s64, err: ", err)
	}
	fmt.Println("Decoded String: ", string(d64))

}
