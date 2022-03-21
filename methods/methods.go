package main

import "fmt"

type user struct {
	name string
	age  uint8
}

func (u *user) fazerNiver() {
	u.age++
}

func (u user) salvar() {
	fmt.Println("the save method was called")
	fmt.Printf("saving %s \n", u.name)
}

func main() {
	user1 := user{"user 1", 20}
	user1.salvar()
	fmt.Println(user1)
	user1.fazerNiver()
	fmt.Println(user1)
}
