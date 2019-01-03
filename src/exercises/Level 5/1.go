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

	fmt.Println(p1.first, p1.last)
	for i, v := range p1.favoriteFlavors {
		fmt.Println(i, v)
	}

	fmt.Println(p2.first, p2.last)
	for i, v := range p2.favoriteFlavors {
		fmt.Println(i, v)
	}
}
