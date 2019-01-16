package main

import (
	"encoding/json"
	"fmt"
)

type edad int
type Persona struct {
	 First, last string
	 age edad
}

func (p Persona) toJson () (string, error) {
	bs, err := json.Marshal(p)

	return string(bs), err
}

func main() {
	ivan := Persona{
		First: "Ivan",
		last: "Ruiz",
		age: 22,
	}

	str, err := ivan.toJson()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf(str)
}
