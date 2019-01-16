package repasoTren

import (
	"encoding/json"
	"fmt"
)

type Persona struct {
	First, Last string
	Age int
}

func main() {
	p1 := Persona{"Ivan", "Ruiz", 22}

	bs, err := json.Marshal(p1)
	if err != nil {
		println(err)
	}
	fmt.Println(string(bs))

	var bs2 Persona

	err = json.Unmarshal(bs, &bs2)
	if err != nil {
		println(err)
	}

	fmt.Print(bs2)
}
