package main



func main() {
	canal := multiplexar(escrever("Olá Mundo!"), escrever("Programando em Go!"))
	for i := 0; i < 10; i++ {
		println(<-canal)
	}

}

func multiplexar(entrada1, entrada2 <-chan string) <-chan string {
	canalSaida := make(chan string)
	go func() {
		for {
			select {
			case mensagem := <-entrada1:
				canalSaida <- mensagem
			case mensagem := <-entrada2:
				canalSaida <- mensagem
			}
		}
	}()
	return canalSaida
}

func escrever(texto string) <-chan string {
	canal := make(chan string)
	go func() {
		for {
			canal <- texto
		}
	}()
	
	return canal
}

