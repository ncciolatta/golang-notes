package main

import "fmt"

func main() {
	retorno := func(text string) string {
		return fmt.Sprintf("Recebido: %s", text)
	}("string parameter") //this sinalizes it

	fmt.Println(retorno)
}
