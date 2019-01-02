package main

import "fmt"

const (
	a = 2019 + iota
	b = a + iota
	c = a + iota
	d = a + iota
	e = a + iota
)

func main() {
	fmt.Println(a, b, c, d, e)
}

