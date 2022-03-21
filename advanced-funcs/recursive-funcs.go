package main

import "fmt"

func fibo(position uint) uint {
	if position <= 1 {
		return position
	}

	return fibo(position-2) + fibo(position-1)
}

func main() {
	//fibonacci

	position_u := uint(10)
	fmt.Println(fibo(position_u))
}
