package main

//funções recursivas

func fibonacci(posicao uint) uint {
	if posicao <= 1 {


		return posicao
	}
	return fibonacci(posicao-2) + fibonacci(posicao-1)
}

// func main() {
// 	posicao := uint(8)
// 	for i:= uint(1); i <= posicao; i++ {
// 		println("fibonacci", fibonacci(i))
// 	}
// }


func main() {
	tarefas := make(chan uint, 45)
	resultados := make(chan uint, 45)

	go worker(tarefas, resultados)
	go worker(tarefas, resultados)
	go worker(tarefas, resultados)
	go worker(tarefas, resultados)
	go worker(tarefas, resultados)
	// são 5 processos que vão ficar esperando por tarefas, e 5 processos que vão ficar esperando por resultados

	for i :=0; i<45; i++{
		tarefas <- uint(i)
	}
	close(tarefas)

	for i :=0 ; i <45; i++ {
		resultado := <- resultados
		println(resultado)
	
	}
}

func worker(tarefas <-chan uint, resultados chan<- uint) {
	for i := range tarefas{
		resultados <- fibonacci(i)
	}
}