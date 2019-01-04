package main

import "fmt"

func main() {
	defer func() {
		fmt.Println("End of program")
	}()

	fmt.Println("Program start")
}



