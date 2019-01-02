package main

import "fmt"

var x int = 42
func main() {

	fmt.Printf("%d\t%b\t%#x", x, x, x)
	x <<= 1
	fmt.Printf("\n%d\t%b\t%#x", x, x, x)

}
