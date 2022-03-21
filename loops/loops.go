package main

import "fmt"

func main() {
	// var i int8 = 0
	// for i < 10 {
	// 	i++
	// 	fmt.Println("incrementando i", i)
	// 	time.Sleep(time.Second)
	// }
	// fmt.Println(i)

	// for y := 0; y < 10; y += 2 {
	// 	fmt.Println("incrementando y", y)
	// 	time.Sleep(time.Second)
	// }

	nomes := [3]string{"João", "Davi", "Lucas"}
	// for i := 0; i < len(nomes); i++ {
	// 	fmt.Println(nomes[i])
	// }

	for i, item := range nomes {
		fmt.Println(i, item)
	}
}
