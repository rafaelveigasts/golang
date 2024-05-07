package main

import (
	"fmt"
	"time"
)

func main() {
	// go escrever("Olá Mundo") // go routine
	// escrever("Programando em Go!")
	go escrever("Olá Mundo")
	escrever("Programando em Go!")
}


func escrever(texto string){
	for {
		fmt.Println(texto)
	time.Sleep(time.Second)
	}

}
// concorrência é !== de paralelismo (concorrência é a capacidade de lidar com muitas coisas ao mesmo tempo), paralelismo é a execução simultânea de processos.

// go routine é um método que permite que uma função seja executada de forma concorrente, ou seja, em paralelo com a função principal.