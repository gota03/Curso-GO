package main

import (
	"fmt"
)

type ID int

var id ID = 1
var nome string = "Mateus"

func main() {
	fmt.Printf("Tipo da variável: %T", id) // %T mostra o tipo de uma variável
	fmt.Printf("Valor: %v", nome)          // %v mostra o valor da variável
}
