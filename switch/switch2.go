package main

import "fmt"

func dayOfWeek(numero int) string {
	switch {
	case numero == 1:
		return "Domingo"
	case numero == 2:
		return "Segunda"
	default:
		return "Número inválido"
	}

}

func main() {
	fmt.Println("switches")
	fmt.Println(dayOfWeek(2), dayOfWeek(3))

}
