package main
import "fmt"

type x interface {}

func main() {

	var x interface{} = 10
	var y interface{} = "Hello, Word!"

	showType(x)
	showType(y)
}

func showType(t interface{})  {
	fmt.Printf("O tipo da varivale é %t eo valor é %v \n", t, t)
}