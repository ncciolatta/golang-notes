package main

import "fmt"

func calculosMatematicos(n1, n2 int) (soma int, subtrac int) {
	soma = n1 + n2
	subtrac = n1 - n2
	return
}

func main() {
	fmt.Println(calculosMatematicos(4, 8))
}
