package main
import "fmt"

func main(){
	salarios := map[string]int{"joao": 30000, "Agata": 10000, "Amabile": 2000}
	fmt.Println(salarios["joao"])

	for _, salario := range salarios{
		fmt.Printf("o Salario de %d \n", salario)
	}
}
