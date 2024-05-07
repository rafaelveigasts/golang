package main

import (
	ftm "fmt"
)

type retangulo struct {
	altura float64
	largura float64
}

type circulo struct {
	raio float64
}

type forma interface { // somente assinaturas de métodos
	area() float64 // método area que retorna um float64
}

func escreverArea(f forma){
	ftm.Printf("A área da forma é %0.2f \n", f.area())
}

func (r retangulo) area() float64 {
	return r.altura * r.largura
}

func (c circulo) area() float64 {
	return 3.14 * (c.raio * c.raio)
}

func main(){
	r := retangulo{10, 15}
	escreverArea(r)

	c := circulo{10}
	escreverArea(c)
}
