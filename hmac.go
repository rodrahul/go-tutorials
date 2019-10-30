package main

import (
	"crypto/hmac"
	"crypto/sha256"
	"fmt"
)

func getHMACCode(s string) string {
	h := hmac.New(sha256.New, []byte(s))
	return fmt.Sprintf("%x", h.Sum(nil))
}

func main() {
	s := "Rahul"
	sHmac := getHMACCode(s)
	fmt.Println(sHmac)

	s = "Rahul"
	sHmac = getHMACCode(s)
	fmt.Println(sHmac)
}