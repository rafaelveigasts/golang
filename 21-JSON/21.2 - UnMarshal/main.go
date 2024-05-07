package main

import (
	"encoding/json"
	"log"
)

type cachorro struct {
	Nome  string `json:"nome"`
	Raça  string `json:"raça"`
	Idade int    `json:"idade"`
}

func main() {
	//json.Marshal() // transforma um struct em json
	//json.Unmarshal() // transforma um json em struct

	cachorroEmJSON := `{"nome":"Rex","raça":"Dálmata","idade":3}`
	var c cachorro
	
	// json.Unmarshal([]byte(cachorroEmJSON), &c)

	if erro := json.Unmarshal([]byte(cachorroEmJSON), &c); erro != nil {
		log.Fatal(erro)
	}

	log.Println(c)

	//converter para map

	cachorroEmJSON2 := `{"nome":"Toby","raça":"Poodle"}`

	c2 := make(map[string]string)

	if erro := json.Unmarshal([]byte(cachorroEmJSON2), &c2); erro != nil {
		log.Fatal(erro)
	}

	log.Println(c2)
}