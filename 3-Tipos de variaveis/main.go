package main

import (
	"errors"
	"fmt"
)

func main() {
	var numero int16 = 100
	fmt.Println(numero)

	//uint // unsigned int int sem sinal

	var numero2 uint32 = 1000
	fmt.Println(numero2)

	var numero3 rune = 12345 // representa um mapeamento da tabela unicode (int32)
	fmt.Println(numero3)

	var erro error = errors.New("Erro interno")
	fmt.Println(erro)
}