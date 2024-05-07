package main

func alunoEstaAprovado(n1, n2 float32) bool {
	defer recuperarExecucao()
	defer println("Média calculada. Resultado será retornado!")
	println("Entrando na função para verificar se o aluno foi aprovado")
	media := (n1 + n2) / 2
	if media >= 6 {
		return true
	}
	panic("A média é inferior a 6!") // panic interrompe a execução do programa
}

func recuperarExecucao() {
	if r := recover(); r != nil {
		println("Tivemos um erro e o programa foi recuperado")
	}
}

func main() {
	println(alunoEstaAprovado(6, 6))
	println("Pós execução")

	println(alunoEstaAprovado(1,2))
}

// panic é diferente do erro do js, ele para a execução do programa, por isso usamos o recover para tratar o erro. Antes dele matar todo o programa ele chama todas as funções que tem defer.