package main

func main() {
	for i := 0; i < 10; i++{
		println(i)
	}

	numeros := [] string{"um", "dois", "três"}
	for k, v := range numeros{
		println(k, v)
	}

	j := 0
	for j < 10{
		println(j)
		j++
	}
}
