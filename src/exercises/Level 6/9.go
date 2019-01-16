package main

import "fmt"

func main() {
	xi := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	fmt.Println(sum(xi...))
	fmt.Println(odd(func (ii ...int) int {
		var total int
		for _, v := range ii {
		total += v
	}

		return total
	}, xi...))
	fmt.Println(even(sum, xi...))
}

func sum(ii ...int) int {
	var total int
	for _, v := range ii {
		total += v
	}

	return total
}

func odd(f func(...int) int, ii ...int) int {
	var odds []int

	for _, v := range ii {
		if v%2 == 1 {
			odds = append(odds, v)
		}
	}

	return f(odds...)
}


func even(f func(...int) int, ii ...int) int {
	var odds []int

	for _, v := range ii {
		if v%2 == 0 {
			odds = append(odds, v)
		}
	}

	return f(odds...)
}
