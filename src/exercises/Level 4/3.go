package main

import "fmt"

func main() {
	x := []int{11,22,33,44,55,66,77,88,99,1010}

	fmt.Println(x[:5])
	fmt.Println(x[5:])
	fmt.Println(x[2:7])
	fmt.Println(x[1:6])
	fmt.Println(x[:])
}
