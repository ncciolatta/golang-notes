package main

import "fmt"

func main() {
	var variavel1 string = "Variavel 1"
	variavel2 := "Variavel 2"
	fmt.Println(variavel1)
	fmt.Println(variavel2)

	var (
		variavel3 string = "dadasdasd"
		variavel4 string = "hfjgjghjgh"
	)
	fmt.Println(variavel3, variavel4)
	variavel5, variavel6 := "var5", "var6"
	fmt.Println(variavel5, variavel6)

	const constant1 string = "const"
	fmt.Println(constant1)

	variavel5, variavel6 = variavel6, variavel5
	fmt.Println(variavel5, variavel6)
}
