package main

import "fmt"

func main() {
	fav := map[string][]string{
		`bond_james`:      []string{`Shaken, not stirred`, `Martinis`, `Women`},
		`moneypenny_miss`: []string{`James Bond`, `Literature`, `Computer Science`},
		`no_dr`:           []string{`Being evil`, `Ice cream`, `Sunsets`},
	}

	fav["ruiz_ivan"] = []string{"esto", "aquello", "lo otro"}

	for k, v := range fav {
		fmt.Printf("Key %v, value -> %v\n", k, v)
		for i2, v2 := range v {
			fmt.Printf("\tIndex %v value -> %v\n", i2, v2)
		}
	}
}
