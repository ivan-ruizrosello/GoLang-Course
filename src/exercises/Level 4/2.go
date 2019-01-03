package main

import "fmt"

func main() {
	x := []int{11,22,33,44,55,66,77,88,99,1010}

	for v, i := range x {
		fmt.Println(v, i)
	}

	fmt.Printf("%T", x)
}
