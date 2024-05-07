package main



func main() {
	canal := escrever("Olá Mundo!")
	for i := 0; i < 10; i++ {
		println(<-canal)
	}
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

// padrao generator é bom para esconder a complexidade