package main

import "fmt"

func diaDaSemana(numero int) string {
	diasDaSemana := []string{"Domingo", "Segunda-feira", "Terça-feira", "Quarta-feira", "Quinta-feira", "Sexta-feira", "Sábado"}

	if numero >= 1 && numero <= 7 {
		return diasDaSemana[numero-1]
	}

	return "Número inválido"
}

func main() {
	fmt.Println(diaDaSemana(8))
}


// fallthrough: é uma palavra-chave que força a execução do próximo case, mesmo que o case atual seja falso.
