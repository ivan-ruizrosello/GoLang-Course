package main

func main() {
	switch {
	case (1==1):
		println("1 == 1 -> true")
		fallthrough
	default:
		println("Default option")
	}
}
