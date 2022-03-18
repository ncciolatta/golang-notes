package main

import "fmt"

type user struct {
	name    string
	age     uint8
	address address
}
type address struct {
	street string
	number uint8
}

func main() {
	fmt.Println("Arquivo de structs")
	addressExample := address{"Willam dasdasd", 245}
	var u user
	u.name = "Davi"
	u.age = 21
	u.address = addressExample
	fmt.Println(u)

	u2 := user{"Jay", 21, addressExample}
	fmt.Println(u2)

	u3 := user{name: "Pedro"}
	fmt.Println(u3)

}
