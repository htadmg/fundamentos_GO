package main

import "fmt"

type x interface{}

func main() {
	var minhaVar interface{} = "Hello, Word!"

	println(minhaVar.(string))
	res, ok := minhaVar.(int)
	fmt.Printf("O valor de res é %v e o resultado de ok é %v", res, ok)
}
