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
	fmt.Printf("o tipo de E é %t", f)
}
