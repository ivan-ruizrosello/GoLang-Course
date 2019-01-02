package main

import "fmt"

var x int = 42
var y string = "James Bond"
var z bool = false

func main() {

	s := fmt.Sprintf("%v\t%v\t%v", x, y, z)
	fmt.Println(s)
	// Se llaman "zero values"
}
