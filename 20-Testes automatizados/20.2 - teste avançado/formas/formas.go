package formas


type Retangulo struct {
	altura float64
	largura float64
}

type Circulo struct {
	raio float64
}

type Forma interface { // somente assinaturas de métodos
	area() float64 // método area que retorna um float64
}

func (r Retangulo) Area() float64 {
	return r.altura * r.largura
}

func (c Circulo) Area() float64 {
	return 3.14 * (c.raio * c.raio)
}
