package main

import "fmt"

func main() {
	var nome interface{} = "Mateus"
	res, stat := nome.(int)
	// ao usar duas variáveis ao tentar converter um tipo, a primeira variável recebe o valor e a segunda recebe true caso dê certo ou false caso dê errado a conversão
	fmt.Printf("O valor de res é: %v\nE o stat: é %t", res, stat)
}
