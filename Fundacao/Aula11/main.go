package main

import "fmt"

type Empresa struct {
	Situacao bool
}

type CNPJ interface {
	DesativarEmpresa()
	// As interfaces são utilizadas para definir comportamentos que tipos concretos (structs, por exemplo) devem implementar. Uma interface define um conjunto de métodos, mas não a implementação deles. Qualquer tipo que implemente esses métodos é considerado uma implementação daquela interface.
	// A relação entre a interface e o tipo concreto se dá apenas pela presença dos métodos necessários
}

func (e Empresa) DesativarEmpresa() {
	// A implementação de uma interface por um tipo é definida unicamente pela assinatura dos métodos. Se um tipo (como uma struct) possui métodos cuja assinatura corresponde aos métodos definidos por uma interface, então esse tipo automaticamente implementa essa interface.

	e.Situacao = true
	fmt.Printf("Situação: %t", e.Situacao)
}

func DesativacaoEmpresa(cnpj CNPJ) {
	// A função DesativacaoEmpresa aceita qualquer tipo que implemente a interface CNPJ. Dessa forma, qualquer struct ou tipo que tenha um método DesativarEmpresa pode ser passado como argumento para essa função, tornando-a mais genérica e flexível.

	cnpj.DesativarEmpresa()
}

func main() {
	empresa := Empresa{}
	DesativacaoEmpresa(empresa)
}
