package main

import "fmt"

func dayOfWeek(numero int) string {
	switch numero {
	case 1:
		return "Domingo"
	case 2:
		return "Segunda"
	default:
		return "Número inválido"
	}

}

func main() {
	fmt.Println("switches")
	fmt.Println(dayOfWeek(2), dayOfWeek(3))

}
