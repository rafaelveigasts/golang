package main

import (
	"bytes"
	"encoding/json"
	"fmt"
)

type cachorro struct {
	Nome  string `json:"nome"`
	Raça  string `json:"raça"`
	Idade int    `json:"idade"`
}

func main() {
	//json.Marshal() // transforma um struct em json
	//json.Unmarshal() // transforma um json em struct

	c := cachorro{"Rex", "Dálmata", 3}
	fmt.Println(c)

	// Transformando o struct em json
	jsonCachorro, erro := json.Marshal(c)
	if erro != nil {
		fmt.Println(erro)
	}
	fmt.Println(jsonCachorro) // retorna um slice de bytes

	// Transformando o slice de bytes em json
	bytes := bytes.NewBuffer(jsonCachorro)
	fmt.Println(bytes)


	c2 := map[string]string{
		"nome": "Toby",
		"raça": "Poodle",
	}

	jsonCachorro2, erro := json.Marshal(c2)
	if erro != nil {
		fmt.Println(erro)
	}

	fmt.Println(string(jsonCachorro2)) // retorna um slice de bytes


	

}