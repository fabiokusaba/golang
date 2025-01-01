package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
)

type cachorro struct {
	Nome  string `json:"nome"`
	Raca  string `json:"raca"`
	Idade uint   `json:"idade"`
}

func main() {
	c := cachorro{"Rex", "Dalmata", 3}
	fmt.Println(c)

	// Convertendo para JSON
	cachorroEmJSON, erro := json.Marshal(c)
	if erro != nil {
		log.Fatal(erro)
	}
	fmt.Println(cachorroEmJSON)

	// Cria um novo Buffer baseado num slice de bytes
	fmt.Println(bytes.NewBuffer(cachorroEmJSON))

	// Criando um cachorro de uma outra maneira sem utilizar struct, mas sim utilizando map
	c2 := map[string]string{
		"nome": "Toby",
		"raca": "Poodle",
	}

	// Convertendo para JSON
	cachorro2EmJSON, erro := json.Marshal(c2)
	if erro != nil {
		log.Fatal(erro)
	}
	fmt.Println(cachorro2EmJSON)

	fmt.Println(bytes.NewBuffer(cachorro2EmJSON))
}
