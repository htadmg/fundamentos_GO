package main

import "fmt"

type MyNumber int

type Number interface{
	~int | float64
}

func soma[T Number](m map[string] T) T{
	var soma T

	for _, v := range m {
		soma += v
	}
	return soma
}

func main() {
	m := map[string] int{"Agata":1000, "Joao": 2000, "Maria":2000}
	m2 := map[string] float64{"Agata": 1000.20, "Joao": 2000.3, "Maria": 2000.6}
	m3 := map[string] MyNumber{"Agata":1000, "Joao": 2000, "Maria":2000}

	fmt.Println(soma(m))
	fmt.Println(soma(m2))
	fmt.Println(soma(m3))
}
