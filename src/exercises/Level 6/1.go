package main

import "fmt"

func main() {
	a := foo()

	b, c := bar()

	fmt.Println(a, b, c)
}

func foo() int {
	return 10
}

func bar() (int, string) {
	return 20, "xd"
}
