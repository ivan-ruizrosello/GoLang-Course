package main

import "fmt"

func main() {

	sl := []int{42, 43, 44, 45, 46, 47, 48, 49, 50, 51}

	fmt.Println(sl)
	sl = append(sl[:3], sl[6:]...)

	fmt.Println(sl)
}
