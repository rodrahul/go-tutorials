package main

import (
	"fmt"
	"net"
)

func main() {

	addr, err := net.LookupHost("rahulrod.com")
	if err, ok := err.(*net.DNSError); ok {
		fmt.Println("Printing out DNSError values")
		fmt.Printf("Err: %s\n", err.Err)
		fmt.Printf("Name: %s\n", err.Name)
		fmt.Printf("Server: %s\n", err.Server)
		fmt.Printf("isTimeout: %v\n", err.IsTimeout)
		fmt.Printf("isTemporary: %v\n", err.IsTemporary)

		// Using underlying methods of dnserror
		if err.Timeout() {
			fmt.Println("operation timed out")
		} else if err.Temporary() {
			fmt.Println("temporary error")
		} else {
			fmt.Println("generic error: ", err)
		}
		return
	}

	fmt.Println("Addr: ", addr)

}

/*
https://godoc.org/net#DNSError
type DNSError
type DNSError struct {
    Err         string // description of the error
    Name        string // name looked for
    Server      string // server used
    IsTimeout   bool   // if true, timed out; not all timeouts set this
    IsTemporary bool   // if true, error is temporary; not all errors set this
}

*DNSError implements the error interface

type DNSError
	func (e *DNSError) Error() string
	func (e *DNSError) Temporary() bool
	func (e *DNSError) Timeout() bool

Once we have the type DNSError we can access the underlying methods
*/
