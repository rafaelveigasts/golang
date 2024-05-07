package formas

import "testing"

func TestArea(t *testing.T) {
	t.Run("Retângulo", func(t *testing.T) {
		retangulo := Retangulo{10, 15}
		areaEsperada := float64(150)
		areaRecebida := retangulo.Area()

		if areaRecebida != areaEsperada {
			t.Errorf("A área recebida %f é diferente da esperada %f", areaRecebida, areaEsperada)
		}
	})

	t.Run("Círculo", func(t *testing.T) {
		circulo := Circulo{10}
		areaEsperada := float64(314)
		areaRecebida := circulo.Area()

		if areaRecebida != areaEsperada {
			t.Errorf("A área recebida %f é diferente da esperada %f", areaRecebida, areaEsperada)
		}
	})
}