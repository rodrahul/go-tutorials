package main

import (
	"encoding/json"
	"fmt"
)

type person struct {
	First string
	Last  string
	Age   int
}

func main() {
	p1 := person{
		First: "James",
		Last:  "Bond",
		Age:   32,
	}

	p2 := person{
		First: "Miss",
		Last:  "Moneypenny",
		Age:   32,
	}

	people := []person{p1, p2}

	fmt.Printf("%+v\n", people)

	bs, err := json.Marshal(people)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(bs))
}

/*
Output is
[{"First":"James","Last":"Bond","Age":32},{"First":"Miss","Last":"Moneypenny","Age":32}]
*/
