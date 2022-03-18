package main

import "fmt"

func main() {
	sum := 1 + 2
	subtrac := 1 - 2
	divs := 10 / 4
	mults := 10 * 5
	modrest := 10 % 2

	fmt.Println(sum, subtrac, divs, mults, modrest)

	var numero1 int16 = 10
	// they arent the same TYPE
	// var numero2 int32 = 25
	var numero2 int16 = 10
	soma := numero1 + numero2
	fmt.Println(soma)

	// logic operators
	//&&
	//||
	//! is negate, invert

	//unary
	//num++
	//num--
	//num +=
	//num -=

}
