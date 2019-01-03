package main

import "fmt"

type Person struct {
	first string
	last string
	favoriteFlavors []string
}

func main() {
	p1 := Person{
		first: "Ivan",
		last: "Ruiz",
		favoriteFlavors: []string{"Vanilla", "Chocolate"},
	}

	p2 := Person{
		first: "Xavi",
		last: "Rodriguez",
		favoriteFlavors: []string{"Fresa", "Menta"},
	}

	personas := map[string]Person{
		p1.last: p1,
		p2.last: p2,
	}

	for k, v := range personas[p1.last].favoriteFlavors {
		fmt.Println(k, v)
	}

	for k, v := range personas {
		println(k)
		fmt.Println(v.first, v.last)
		for i, val := range v.favoriteFlavors {
			println(i, val)
		}
	}
}
