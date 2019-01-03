package main

import "fmt"

type vehicle struct {
	doors int
	color string
}

type truck struct {
	vehicle
	fourWheel bool
}

type sedan struct {
	vehicle
	luxury bool
}

func main() {
	s := sedan{
		vehicle: vehicle{
			doors: 4,
			color: "green",
		},
		luxury: false,
	}

	t := truck{
		vehicle: vehicle{
			doors: 4,
			color: "green",
		},
		fourWheel: true,
	}

	fmt.Println(t)
	fmt.Println(s)



	fmt.Println(t.vehicle.color)
	fmt.Println(s.doors)
}
