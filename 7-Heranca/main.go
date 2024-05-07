package main

func main() {
	println("Herança")

	type endereco struct {
		rua  string
		num  int
	}

	type person struct {
		name string
		age  int
		nice bool
		endereco 
	}


	var u person
	u.name = "John"
	u.age = 30
	println(u.name, u.age)

	u2 := person{name: "Jane", age: 25, nice: true}
	println(u2.name, u2.age, u2.nice)

	u3 := person{"Doe", 40, false, endereco{"Rua 1", 123}}
	println(u3.name, u3.age, u3.nice, u3.endereco.rua, u3.endereco.num)
}