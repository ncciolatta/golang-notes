package main

import (
	"encoding/json"
	"fmt"
	"log"
)

type cachorro struct {
	Nome  string `json:"nome"`
	raca  string `json:"raca"`
	idade uint   `json:"age"`
}

func main() {
	cachorroEmJSON := `{"nome":"Maru","raca":"Pitbull", "age":3}`

	var c cachorro

	if erro := json.Unmarshal([]byte(cachorroEmJSON), &c); erro != nil {
		log.Fatal(erro)
	}
	fmt.Println(c)
	cachorro2EmJSON := `{"nome":"jao","raca":"gatochorro"}`

	c2 := make(map[string]string)
	if erro := json.Unmarshal([]byte(cachorro2EmJSON), &c2); erro != nil {
		log.Fatal(erro)
	}
	fmt.Println(c2)
}
