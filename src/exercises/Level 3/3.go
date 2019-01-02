package main

func main() {
	born := 1996
	actualYear := 2019

	println("Years i've been alive:")

	for born <= actualYear {
		println(born)
		born ++
	}
}
