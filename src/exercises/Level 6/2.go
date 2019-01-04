package main

import "fmt"

func main() {
	arr := []int{1,2,3,10}

	res := foo(arr...)
	res2 := bar(arr)

	fmt.Println(res)
	fmt.Println(res2)
}

func foo(x ...int) int {
	fmt.Println(x)

	var sum int
	for _, v := range x {
		sum += v
	}

	return sum
}

func bar(x []int) int {
	var sum int
	for _, v := range x {
		sum += v
	}

	return sum
}