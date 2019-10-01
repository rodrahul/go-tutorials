package main

import (
	"fmt"
	"os"
	"path/filepath"
)

func main() {
	p := filepath.Join("dir1", "dir2", "filename")
	fmt.Println("p:", p)
	p = os.Args[0]
	fmt.Println(os.Args)

	fmt.Println(filepath.Join("dir1//", "filename"))
	fmt.Println(filepath.Join("dir1/../dir1", "filename"))

	fmt.Println("Dir(p):", filepath.Dir(p))
	fmt.Println("Base(p):", filepath.Base(p))

	fmt.Println(filepath.IsAbs("a/b"))
	fmt.Println(filepath.IsAbs("/a/b"))
}
