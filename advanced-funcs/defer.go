package main

import "fmt"

func function1() {
	fmt.Println("execution of func 1")

}
func function2() {
	fmt.Println("execution of func 2")
}

func main() {
	defer function1()
	function2()
}
