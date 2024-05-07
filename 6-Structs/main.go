package main

type person struct {
	name string
	age  int
	nice bool
	endereco 
}

type endereco struct {
	rua  string
	num  int
}

func main() {
	println("Structs")
	var u person
	u.name = "John"
	u.age = 30
	println(u.name, u.age)

	u2 := person{name: "Jane", age: 25, nice: true}
	println(u2.name, u2.age, u2.nice)

	u3 := person{"Doe", 40, false, endereco{"Rua 1", 123}}
	println(u3.name, u3.age, u3.nice, u3.endereco.rua, u3.endereco.num)

}

// structs são tipos de dados que podem ser usados para criar estruturas de dados mais complexas e personalizadas.
// Eles são compostos por um conjunto de campos, que podem ser de diferentes tipos de dados.
// A declaração de um struct é feita com a palavra-chave type seguida do nome do struct e a palavra-chave struct.
