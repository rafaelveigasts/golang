package main

type usuario struct {
	nome string
	idade uint8
}

func (u usuario) salvar() { // aqui está o detalhe do método, a função salvar agora é um método do tipo usuario, o "u" é o this do js
	println("Salvando os dados do usuário", u.nome)


}

func (u usuario) maiorDeIdade() bool {
	return u.idade >= 18
}

func (u *usuario) fazerAniversario() {
	u.idade++
}

func main(){
	u := usuario{"Fulano", 33}
	u.salvar()
	println(u.maiorDeIdade())
	println(u.idade)
	u.fazerAniversario()
	println(u.idade)
}
