package main

import "fmt"

func main() {
	var variavel1 string = "Variavel 1" // explíocito
	fmt.Println(variavel1)

	variavel2 := "Variavel 2" // implicito
	fmt.Println(variavel2)

	variavel3, variavel4 := "Variavel 3", "Variavel 4" // multiplas variaveis
	fmt.Println(variavel3, variavel4)
	
}