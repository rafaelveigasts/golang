package main

// pense nos ponteiros(referência de memória) como uma forma de acessar a memória de um valor

func main() {
	println("Ponteiros")

	var i int = 10
	var p *int = &i // p é um ponteiro para i, o & é usado para obter o endereço de memória de i

	println(i, p, *p)

	*p = 20 // altera o valor de i
	println(i, p, *p) // *p é a desreferência de p, ou seja, o valor de i

}

// A vantagem de usar ponteiros é que podemos passar referências de memória para funções, o que pode ser mais eficiente do que passar valores. 