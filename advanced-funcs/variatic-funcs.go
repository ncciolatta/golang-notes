package main

import "fmt"

func sum(numbers ...int) int { //becomes a slice
	total := 0
	for _, num := range numbers {
		total += num
	}
	return total

}
func main() {
	fmt.Println(sum(5, 5, 5, 5, 5, 5, 5, 5))
}
