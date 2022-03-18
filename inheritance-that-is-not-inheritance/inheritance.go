package main

import "fmt"

type person struct {
	name   string
	age    uint8
	height uint8
}

type student struct {
	person
	college    string
	graduation string
}

func main() {
	fmt.Println("Inheritance")
	person1 := person{"James", 21, 192}
	fmt.Println(person1)

	student1 := student{person1, "University of Miami", "CS"}
	fmt.Println(student1)
	fmt.Println(student1.name)
}
