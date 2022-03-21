package main

import "fmt"

func revertSignal(num int) int {
	return num * -1
}

func revertSignalPointerv(num *int) {
	*num = *num * -1
}

func main() {
	num := 20
	revertedNum := revertSignal(num)
	fmt.Println(revertedNum)

	newNum := 40
	fmt.Println(newNum)
	revertSignalPointerv(&newNum)
	fmt.Println(newNum)
}
