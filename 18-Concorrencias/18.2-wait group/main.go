package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	var waitGroup sync.WaitGroup // cria um grupo de espera
	waitGroup.Add(2) // adiciona 2 funções ao grupo de espera
	

	go func(){
		escrever("Olá Mundo")
		waitGroup.Done() // avisa que a função terminou, passa -1 para o Add
	}()

	go func(){
		escrever("Programando em Go!")
		waitGroup.Done() // avisa que a função terminou
	}()

	waitGroup.Wait() // espera todas as funções terminarem e que o add chegue a 0
}


func escrever(texto string){
	for i:=0; i<5; i++{
		fmt.Println(texto)
	time.Sleep(time.Second)
	}

}
// concorrência é !== de paralelismo (concorrência é a capacidade de lidar com muitas coisas ao mesmo tempo), paralelismo é a execução simultânea de processos.

// go routine é um método que permite que uma função seja executada de forma concorrente, ou seja, em paralelo com a função principal.