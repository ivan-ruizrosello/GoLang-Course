package main

import "fmt"

func main() {
	fmt.Println("Hello")

	foo()

	fmt.Println("Something more")
	for i := 0; i < 100; i++ {
		if i%10 == 0 {
			fmt.Println(i)
		}
	}
}

func foo() {
	fmt.Println("Im in foo")
}
