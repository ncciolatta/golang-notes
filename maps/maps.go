package main

import "fmt"

func main() {
	fmt.Println("Maps")
	usuario := map[string]string{ //this map/dict has string keys and string values
		"nome":      "Pedro",
		"sobrenome": "Silva",
	}

	fmt.Println(usuario["nome"])

	usuario2 := map[string]map[string]string{ // this map/dict has a map key carrying string keys and string keys
		"nome": {
			"primeiro": "João",
			"ultimo":   "Pedro",
		},
		"curso": {
			"nome":   "Engenharia",
			"campus": "Campus 1",
		},
	}
	fmt.Println(usuario2)
	delete(usuario2, "nome")
	fmt.Println(usuario2)

	usuario2["signo"] = map[string]string{
		"nome": "Câncer",
	}
	fmt.Println(usuario2)
}
