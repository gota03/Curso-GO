package main

import "fmt"

type Cliente struct {
	Nome  string
	Idade int
	Ativo bool
	Endereco
}

type Endereco struct {
	Estado string
	Cidade string
	Rua    string
}

func (cliente Cliente) DesativarCliente() {
	//(cliente Cliente) define que o método pertence a struct Cliente, possui o nome de receiver
	//O receiver pode ser visto como um "argumento especial" que sempre está presente quando o método é chamado. Ele indica que esse método pode ser chamado em uma instância de Cliente.
	//Esse receiver (cliente Cliente) é passado por valor, o que significa que qualquer modificação feita em cliente dentro do método não vai alterar a instância original da struct.

	cliente.Ativo = false
	fmt.Printf("Nome do cliente: %s", cliente.Nome)
	fmt.Printf("\nStatus do cliente: %t", cliente.Ativo)
}

// Uma struct (ou estrutura) é uma coleção de campos, ou seja, é um tipo de dado composto que permite agrupar variáveis (chamadas de campos) sob um mesmo nome. Cada campo pode ter um tipo de dado diferente
// Uma struct pode conter outras structs como campos
// Uma struct pode ter funções associadas a ela desde que a função se referencie a struct passando o nome da variável e em seguida o nome da struct

func main() {
	mateus := Cliente{
		Nome:  "Mateus",
		Idade: 21,
		Ativo: true,
		Endereco: Endereco{
			Estado: "Maranhão",
			Cidade: "São Luis",
			Rua:    "14",
		},
	}
	mateus.DesativarCliente()
	//fmt.Printf("Nome: %s \nIdade: %d \nAtivo: %t \nEstado: %s \nCidade %s \nRua: %s", mateus.Nome, mateus.Idade, mateus.Ativo, mateus.Estado, mateus.Cidade, mateus.Rua)
}
