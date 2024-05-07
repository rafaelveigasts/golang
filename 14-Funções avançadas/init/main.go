package main

func init() {
	println("Função Init")
}

func main() {
	println(("Função Main"))
}

//função init roda antes da função main e é possível ter 1 por arquivo e não uma por pacote como a função main, bastante utilizada para setar valores como variáveis de ambiente, conexão com banco de dados, etc.