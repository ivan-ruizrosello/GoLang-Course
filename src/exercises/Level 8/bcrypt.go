package main

import (
	"fmt"
	"golang.org/x/crypto/bcrypt"
)

func main() {
	s := `password123`
	bs, err := bcrypt.GenerateFromPassword([]byte(s), bcrypt.MinCost)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(s)
	fmt.Println(bs)


	err2 := bcrypt.CompareHashAndPassword(bs, []byte(s))
	if err2 != nil {
		println(err2, "You cant login")
		return
	}

	fmt.Println("You're logged in")
}
