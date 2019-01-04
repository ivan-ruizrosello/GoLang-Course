package main

import "fmt"

type person struct {
	first string
	last  string
	age   int
}

func (p *person) changeMe(newPerson person) {
	*p = newPerson
	(*p).first = "JAJA"
}

func main() {
	p := person{
		first: "Ivan",
		last:  "Ruiz",
		age:   22,
	}

	fmt.Println(p)
	p.changeMe(
		person{
			"Xavi", "Rodriguez", 40,
		})

	fmt.Println(p)
}
