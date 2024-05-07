package main

func invertSignal(number *int) {
	*number = *number * -1 // *number é o endereço de memória e o =*numero é o valor referenciado pelo endereço de memória
}

func main() {
	number := 20
	invertSignal(&number)
	println(number) // -20
}