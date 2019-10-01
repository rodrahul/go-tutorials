/*
ioutil.ReadFile will read the file and keep the entire file in the memory
In this program we'll be reading file in chunks,
We'll be doing it using command line parsing
Use flag package to parse commands

usage go run readFile2_golangbot.go -fpath=/path-of-file/test.txt
*/

package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
)

func main() {
	filePath := flag.String("filePath", "readFile.txt", "file path to read file from ")
	flag.Parse()
	fmt.Println("filePath :", *filePath)

	filePtr, err := os.Open(*filePath)
	defer func() {
		if err = filePtr.Close(); err != nil {
			log.Fatal("Error when closing file: ", err)
		}
	}()

	if err != nil {
		fmt.Println("Failed to open file, error: ", err)
		return
	}
	fmt.Println(filePtr)
	fmt.Println("File Name: ", filePtr.Name())

	//--- To read n bytes at a time e.g. 3 below-------------
	//b := make([]byte, 3)
	//for {
	//	_, err := filePtr.Read(b)
	//	if err != nil {
	//		fmt.Println("Error reading file:", err)
	//		break
	//	}
	//	fmt.Println(string(b))
	//}
	//--------------------------------------------------------

	//----- To read line by line -----------------------------
	// buffio.NewScanner needs io.Reader, filePtr implements Read method, so we can pass filePtr into NewScanner
	s := bufio.NewScanner(filePtr)
	for s.Scan() {
		fmt.Println(s.Text())
	}
	err = s.Err()
	if err != nil {
		log.Fatal(err)
	}

	//--------------------------------------------------------

}
