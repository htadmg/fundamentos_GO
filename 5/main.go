package main
import "fmt"

const a = "Hello Word!"
type ID int 

var (
	b bool = true
	c int = 20
	d string = "bla bla"
	e float64 = 1.2
	f ID = 1
)

func main(){
	var meuArray [3] int
	meuArray[0] = 10
	meuArray[1] = 20
	meuArray[2] = 30

	for i, v := range meuArray{
		fmt.Printf("O valor do indice %d é %d \n", i, v)
	}
}
