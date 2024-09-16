package main

import (
	"curso-go/1-Fundacao/Aula18/matematica"
	"fmt"

	"github.com/google/uuid"
)

func main() {
	a := 10
	b := 20
	soma := matematica.Soma(a, b)
	fmt.Printf("Resultado: %v\n", soma)
	println(matematica.A)
	matematica.Mensagem()
	carro := matematica.Carro{
		Marca: "fiat",
	}
	fmt.Printf("\nMarca do carro: %v", carro.Marca)
	fmt.Printf("%v", uuid.New())
}
