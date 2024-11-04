package main

func main() {
	a := 10
	println(a)
	println(&a)

	var ponteiro *int = &a
	println(ponteiro)

	*ponteiro = 20
	println(a)
}