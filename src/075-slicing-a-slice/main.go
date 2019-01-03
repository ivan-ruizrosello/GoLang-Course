package main

import "fmt"

func main() {
	x := []int{11, 22, 33, 44, 55}
	fmt.Println(x)
	fmt.Println(x[:4])
	fmt.Println(x[3:])
	fmt.Println(x[1:3])
}