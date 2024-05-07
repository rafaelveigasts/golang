package main

func main() {
	// for init; condition; post {}
	// for i := 0; i < 10; i++ {
	// 	println("Incrementando", i)
	// }

	// for condition {}
	// j := 0
	// for j < 10 {
	// 	println("Incrementando", j)
	// 	j++
	// }


	// range

	// slice := []string{"banana", "maçã", "laranja"}
	// for i, fruta := range slice {
	// 	println(i, fruta)
	// }

	// for _, fruta := range slice {
	// 	println(fruta)
	// }

	for indice, letra :=range "PALAVRA" {
		println(indice, string(letra))
	}

	usuario :=map[string]string{
		"nome": "Leonardo",
		"sobrenome": "Miranda",
	}

	for chave, valor := range usuario {
		println(chave, valor)
	}

	// // for {}
	// for {
	// 	println("Loop infinito")
	// }
}

//package main

//func main() {}