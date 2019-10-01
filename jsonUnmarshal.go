package main

import (
	"encoding/json"
	"fmt"
)

type person struct {
	First string `json:"First"`
	Last  string `json:"Last"`
	Age   int    `json:"Age"`
}

func main() {
	s := `[{"First":"James","Last":"Bond","Age":32},{"First":"Miss","Last":"Moneypenny","Age":32}]`

	bs := []byte(s)

	fmt.Printf("%T\n", s)
	fmt.Printf("%T\n", bs)

	var people []person

	err := json.Unmarshal(bs, &people)
	if err != nil {
		fmt.Println("error:", err)
	}
	fmt.Printf("%+v\n", people)
}

/*
Output is
[{"First":"James","Last":"Bond","Age":32},{"First":"Miss","Last":"Moneypenny","Age":32}]
*/
