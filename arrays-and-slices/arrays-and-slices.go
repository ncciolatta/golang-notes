package main

import "fmt"

func main() {
	fmt.Println("Arrays and Slices")

	var array1 [5]string
	array1[0] = "Posição 1"
	fmt.Println(array1)

	array2 := [5]string{"dasda", "sdadada", "dadsadasd", "hfjtjjtgjyfg", "gtryhrh"}
	fmt.Println(array2)

	array3 := [...]int{1, 2, 3, 4, 5} //the three dots doesnt make it dinamic, but fixes it to the array size
	fmt.Println(array3)

	slice := []int{11, 12, 13, 14, 15, 16, 1616, 16, 161, 61, 6}
	fmt.Println(slice)

	slice = append(slice, 10)
	fmt.Println(slice)

	slice2 := array2[1:3] //the third is exclusive
	fmt.Println(slice2)

	array2[1] = "new position"
	fmt.Println(slice2)

	// intern arrays
	fmt.Println("----------------")
	slice3 := make([]float64, 10, 11) //(type, elements in it, max capacity of elements)
	fmt.Println(slice3)
	fmt.Println(len(slice3))
	fmt.Println(cap(slice3))

	slice3 = append(slice3, 5)
	slice3 = append(slice3, 6)

	fmt.Println(slice3)
	fmt.Println(len(slice3))
	fmt.Println(cap(slice3))
}
