package main

func funcao1() {
	println("Executando a função 1")
}

func funcao2() {
	println("Executando a função 2")
}

func alunosAprovados(n1, n2 float32) bool {
	defer println("Média calculada. Resultado será retornado!")
	println("Entrando na função para verificar se o aluno foi aprovado")
	media := (n1 + n2) / 2
	if media >= 6 {
		return true
	}
	return false
}

func main() {
	// defer funcao1() //defer adia a execução da função até o final da execução do método main
	// funcao2()
	println(alunosAprovados(7, 8))
}