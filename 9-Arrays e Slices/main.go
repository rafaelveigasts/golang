package main

func main() {
	println("Arrays e Slices")

	var a [5]int
	a[2] = 7
	println(a[0],a[2])

	b := [5]int{1, 2, 3, 4, 5}
	println(b[2])

	c := [...]int{1, 2, 3, 4, 5}
	println(c[2])

	d := [2][3]int{{1, 2, 3}, {4, 5, 6}}
	println(d[1][2])

	e := []int{1, 2, 3, 4, 5}
	println(e[2])

	e = append(e, 6)
	println(e[5])

	f := e[1:3]
	println(f[1])

	g := e[:3]
	println(g[2])

	h := e[1:]
	println(h[2])

	i := make([]int, 5)
	println(i[2])

	j := make([]int, 5, 10) // essa função cria um slice com 5 elementos e capacidade para 10 elementos
	println(j[2])
}