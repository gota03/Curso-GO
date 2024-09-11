package matematica

import "fmt"

func Soma[T int | float64](a, b T) T {
	return a + b
}

var A int = 10

type Carro struct {
	Marca string
}

func Mensagem() {
	fmt.Printf("Uma mensagem")
}

// Para que funções, variáveis, structs, métodos possam ser exportados e visiveís a outros arquivos, precisam começar com a letra maiúscula
