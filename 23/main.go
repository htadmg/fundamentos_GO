package main

import (
	"encoding/json"
	"os"
)

type Conta struct {
	Numero int `json:"numero"` 
	Saldo  int `json:"saldo"`
}

func main() {
	conta := Conta{Numero: 1, Saldo: 100}
	res, err := json.Marshal(conta)
	if err != nil {
		println(err)
	}
	println(string (res))

	encoder := json.NewEncoder(os.Stdout)
	encoder.Encode(conta)

}
