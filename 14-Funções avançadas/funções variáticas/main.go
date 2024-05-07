//função que pode receber n argumentos

package main

func somaNumeros(numeros ...int) int {
	soma := 0
	for _, numero := range numeros {
		soma += numero
	}
	return soma
}

func escrever(texto string, numeros ...int) { // no máximo 1 parâmetro variático por função e obrigatóriamente no final
	for _, numero := range numeros {
		println(texto, numero)
	}
}

func main() {
	println(somaNumeros(1, 2, 3, 4, 5))
	println(somaNumeros())

	escrever("Olá", 1, 2, 3, 4, 5)
}