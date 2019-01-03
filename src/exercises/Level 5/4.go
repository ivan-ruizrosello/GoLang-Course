package main

import "fmt"

func main() {
	s := struct {
		first string
		friends map[string]int
		favDrinks []string
	}{
		first: "Ivan",
		friends: map[string]int{
			"Antonio": 555,
			"Q": 777,
		},
		favDrinks: []string{"Water", "Cola"},
	}

	fmt.Println(s)
}
