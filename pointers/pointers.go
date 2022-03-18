package main

import "fmt"

func main() {
	fmt.Println("Pointers")

	var variavel1 int = 10
	var variavel2 int = variavel1
	variavel1++

	fmt.Println(variavel1, variavel2)

	//pointers are a reference in memory
	var variavel3 int
	var ponteiro *int //int pointer
	variavel3 = 100
	ponteiro = &variavel3
	fmt.Println(variavel3, ponteiro) //dereferencing
	variavel3 += 50
	fmt.Println(variavel3, *ponteiro)
}
