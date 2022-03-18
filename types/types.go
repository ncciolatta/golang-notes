package main

import (
	"errors"
	"fmt"
)

func main() {
	// int8, int16, int32, int64
	// the numbers are bits
	var numero int8 = 100
	fmt.Println(numero)

	// uint does not accept negative
	var numero2 uint32 = 10000
	fmt.Println(numero2)

	//alias
	// INT32 = RUNE
	var numero3 rune = 1234
	fmt.Println(numero3)
	// BYTE = UINT8
	var numero4 byte = 123
	fmt.Println(numero4)
	var numeroReal1 float32 = 123.45
	fmt.Println(numeroReal1)

	var str string = "Texto"
	fmt.Println(str)

	str2 := "Texto2"
	fmt.Println(str2)

	// will print the positing in the ascii table, theres no chars in Go
	char := 'B'
	fmt.Println(char)

	var text string
	fmt.Println(text)

	var boolean bool = true
	fmt.Println(boolean)

	// var erro error
	var erro error = errors.New("error lmao!")
	fmt.Println(erro)
}
