package main

import (
	"fmt"
	"time"
)

func main() {
	canal := make(chan string)	
	go escrever("Olá Mundo", canal)

	// for{

	// 	mensagem := <- canal // canal espera/recebe um valor
	// 	fmt.Println(mensagem)
	//} // isso da um erro de deadlock, pois o canal espera um valor e não tem ninguém para enviar e isso não é perceptível em tempo de compilação, somente em tempo de execução

	for {
		mensagem, aberto := <- canal
		if !aberto {
			break
		}
		fmt.Println(mensagem)
	}

}


func escrever(texto string, canal chan string){
	for i:=0; i<5; i++{
		canal <- texto // atribui o valor ao canal
	time.Sleep(time.Second)
	}

	close(canal) // fecha o canal

}
