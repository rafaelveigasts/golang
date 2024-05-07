package main

import (
	"fmt"
)

func main() {
	canal := make(chan string, 2) // canal com buffer de 2 posições	, isso significa que o canal pode receber 2 valores sem precisar de um receptor

	canal <- "Olá Mundo" // atribui o valor ao canal
	canal <- "Programando em Go!" // atribui o valor ao canal
	//canal <- "teste" // erro de deadlock, pois o canal só pode receber 2 valores

	mensagem := <- canal
	mensagem2 := <- canal
	fmt.Println(mensagem)
	fmt.Println(mensagem2)
	
}

