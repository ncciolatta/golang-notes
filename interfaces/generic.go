package main

import "fmt"

func generica(interf interface{}) {
	fmt.Println(interf)

}

func main() {
	generica("asdads")
	generica(1)
	generica(true)
}
