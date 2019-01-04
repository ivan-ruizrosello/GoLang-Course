package main

import "fmt"

type person struct {
	first string
	last  string
	age   int
}

type secretAgent struct {
	person
	ltk bool
}

func (p person) speak() {
	fmt.Println("I am ", p.first, p.last, " - The person speak")
}

func (p secretAgent) speak() {
	fmt.Println("I am ", p.first, p.last, " - The secretAgent speak")
}

type human interface {
	speak()
}

func bar(h human) {
	var nombre string
	switch h.(type) {
	case person:
		nombre = h.(person).first
	case secretAgent:
		nombre = h.(secretAgent).person.first
	}
	fmt.Println("I was passed into bar", nombre)
}

func main() {
	p1 := person{
		first: "Ivan",
		last:  "Ruiz",
		age:   22,
	}

	sa1 := secretAgent{
		person: person{
			first: "James",
			last:  "Bond",
			age:   34,
		},
		ltk: true,
	}

	sa1.speak()
	p1.speak()
	fmt.Printf("%T\n", sa1)

	bar(p1)
	bar(sa1)
}

