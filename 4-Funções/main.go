package main

func somar(n1 int8, n2 int8) int8 {
	return n1 + n2

}

//funções podem ter mais de um retorno

func calculosMatematicos(n1, n2 int8) (int8, int8) {
	soma := n1 + n2
	subtracao := n1 - n2
	return soma, subtracao
}

func main () {
	resultado := somar(10, 20)
	println(resultado)

	resultadoSoma, resultadoSubtracao := calculosMatematicos(10, 20)
	println(resultadoSoma, resultadoSubtracao)

	// _ ignora o retorno
	resultadoSoma2, _ := calculosMatematicos(10, 20)
	println(resultadoSoma2)
}