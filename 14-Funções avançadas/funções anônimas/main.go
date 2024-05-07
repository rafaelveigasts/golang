package main


func main() {
	//função anônima
	func() {
		println("Olá")
	}()

	//função anônima com parâmetros
	func(texto string) {
		println(texto)
	}("Olá")

	//função anônima com retorno
	retorno := func(texto string) string {
		return texto
	}("Olá")
	println(retorno)

	//função anônima com retorno nomeado
	retornoNomeado := func(texto string) (retorno string) {
		retorno = texto
		return
	}("Olá")
	println(retornoNomeado)
}