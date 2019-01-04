package main

import "fmt"

type person struct {
	first string
	last  string
	age   int
}

func (p person) speak() {
	fmt.Println("Hello, my name is", p.first, p.last, "and im", p.age)
}

func main() {
	p := person{"Ivan", "Ruiz", 22}
	fmt.Println(p)
	p.speak()
}
