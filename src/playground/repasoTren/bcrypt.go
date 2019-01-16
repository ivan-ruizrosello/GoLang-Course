package main

import (
	"fmt"
	"golang.org/x/crypto/bcrypt"
)

func main() {
	pass := "micontra234"

	bs, err := bcrypt.GenerateFromPassword([]byte(pass), bcrypt.MinCost)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(pass, bs)

	err = bcrypt.CompareHashAndPassword(bs, []byte(pass))
	if err != nil {
		fmt.Println("You cant login", err)
		return
	}

	fmt.Println("You are logged in")

}
