package main

func main() {
	born := 1996
	actualYear := 2019

	println("Years i've been alive:")

	for {
		if born > actualYear {
			break
		}

		println(born)
		born ++
	}
}