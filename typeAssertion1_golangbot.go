package main

import (
	"fmt"
	"os"
)

func main() {
	f, err := os.Open("/test.txt")
	if err != nil {
		fmt.Println(err)
		if v, ok := err.(*os.PathError); ok {
			fmt.Println("Op:", v.Op, "path:", v.Path, "err:", v.Err)
		}
		return
	}
	fmt.Println(f.Name(), "opened successfully")
}
/*
If you read the documentation of the Open function carefully, you can see that it returns an error of type *PathError. PathError is a struct type and its implementation in the standard library is as follows,

type PathError struct {
	Op   string
	Path string
	Err  error
}

func (e *PathError) Error() string

From the above code, you can understand that *PathError implements the error interface by declaring the Error() string method

i.e. both *os.PathError and error are of the type error interface and we can do a type assertion

We want to assert the returned error type to type *os.PathError
(Note: *os.PathError implements the interface and not os.PathError)

which can be done as
err.(*os.PathError)

Once we have the type *os.PathError which is a struct we can access the underlying methods attached to the struct type

*/