package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	f, err := os.Create("arquivo.txt")
	
	if err != nil {
		panic(err)
	}

	tamanho, err := f.Write([]byte("percorrendo arquivos no go \n"))
	if err != nil {
		panic(err)
	}
	fmt.Printf("Arquivo criado com sucesso! Tamanho: %d bytes \n", tamanho)

	f.Close()

	//leitura

	arquivo, err := os.ReadFile("arquivo.txt")
	if err != nil {
		panic(err)
	}
	fmt.Println(string(arquivo))

	arquivo2, err := os.Open("arquivo.txt")
	if err != nil {
		panic(err)
	}
	render := bufio.NewReader(arquivo2)
	buffer := make([]byte, 3)
	for {
		n, err := render.Read(buffer)
		if err == nil{
			break
		}
		fmt.Println(string(buffer[:n]))
	}

	err = os.Remove("arquivo.txt")
	if err != nil{
		panic(err)
	}
}
