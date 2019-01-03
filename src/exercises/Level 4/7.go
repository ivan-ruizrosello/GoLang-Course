package main

import "fmt"

func main() {
	a := []string{"James", "Bond", "Shaken, not stirred"}

	b := []string{"Miss", "Moneypenny", "Helloooooo, James."}

	x := [][]string{a, b}

	fmt.Println(x)

	for i, v := range x {
		println("Index:", i)
		for i2, v2 := range v {
			fmt.Printf("\tIndex %v value: %v \n",i2, v2)
		}
	}
}
