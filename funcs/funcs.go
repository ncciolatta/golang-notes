package main

import "fmt"

func sum(num1 uint8, num2 uint8) uint8 {
	return num1 + num2
}

func calcsMath(n1, n2 int8) (int8, int8) {
	soma := n1 + n2
	subtrac := n1 - n2
	return soma, subtrac
}
func main() {
	sumcall := sum(10, 20)
	fmt.Println(sumcall)

	var f = func(txt string) string {
		fmt.Println(txt)
		return txt
	}

	resultado := f("function test!")
	fmt.Println(resultado)

	resultsSoma, resultsSubtrac := calcsMath(10, 15)
	fmt.Println(resultsSoma, resultsSubtrac)
}
