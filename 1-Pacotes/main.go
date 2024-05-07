package main

import (
	"fmt"
	"modulo/auxiliar"

	"github.com/badoux/checkmail"
)

func main() {
	fmt.Println("Olá, mundo! Main")
	auxiliar.Escrever()
	erro := checkmail.ValidateFormat("a@b.c")
	fmt.Println((erro))
	erro2 := checkmail.ValidateFormat("123")
	fmt.Println((erro2))
}
// go build 
// para pegar pacotes externos utilizamos o go get + o link do pacote que queremos geralmente no github
// go mod tidy para atualizar o arquivo go.mod