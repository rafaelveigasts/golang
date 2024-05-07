package main

func main() {
	numero :=10

	if numero > 15 {
		println("maior que 15")
	} else {
		println("menor que 15")
	}

	
//if init; condition : a variável é limitada ao escopo do if
	if numero2 := 10; numero2 > 0 {
		println("maior que 0")
	} else {
		println("menor que 0")
	}
}