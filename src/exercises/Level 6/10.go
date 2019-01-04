package main

import "fmt"

func main() {
	var x, y int

	x, y= 2, 10

	println(x, y)

	{
		temp := x
		x = y
		y = temp
	}

	println(x, y)

	x, y = y, x

	println(x, y)


	f := iniciar()
	f2 := iniciar()

	fmt.Println(f())
	fmt.Println(f())
	fmt.Println(f())
	fmt.Println(f2())
	fmt.Println(f2())
	fmt.Println(f2())
	fmt.Println(f2())
}

func iniciar () func() int{
	var x int

	return func() int {
		x++
		return x
	}
}
