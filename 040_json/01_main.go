package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type person struct {
	Fname string
	Lname string
	Items []string
}

func main() {
	http.HandleFunc("/", foo)
	http.HandleFunc("/mshl", mshl)
	http.HandleFunc("/encd", encd)
	http.ListenAndServe(":8080", nil)
}

func foo(w http.ResponseWriter, req *http.Request) {
	s := `<!DOCTYPE html>
		<html lang="en">
		<head>
		<meta charset="UTF-8">
		<title>FOO</title>
		</head>
		<body>
		You are at foo
		</body>
		</html>`
	w.Write([]byte(s))
}

func mshl(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	p1 := person{
		Fname: "Rahul",
		Lname: "Rode",
		Items: []string{"one", "two", "three"},
	}
	b, err := json.Marshal(p1)
	if err != nil {
		fmt.Println(err)
	}
	w.Write(b)
}

func encd(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	p1 := person{
		Fname: "James",
		Lname: "Bond",
		Items: []string{"Suit", "Gun", "Wry sense of humor"},
	}
	err := json.NewEncoder(w).Encode(p1)
	if err != nil {
		log.Println(err)
	}
}
