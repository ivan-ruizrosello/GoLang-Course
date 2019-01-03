package main

import "fmt"

func main() {
	x := [5]int{11,22,33,44,55}

	for v, i := range x {
		fmt.Println(v, i)
	}

	fmt.Printf("%T", x)
}
