package main

type Conta struct {
	Saldo int
}

func (c *Conta) emprestimo(valor int) int {
	// Ao colocar um * estou dizendo que o endereço de memória onde está armazenado a struct, qualquer atribuição de valores vai ser alterado diretamente na struct e não em uma cópia
	// Antes o valor de c.Saldo só era alterado no escopo local da função emprestimno
	c.Saldo += valor
	println("Valor após empréstimo na função emprestimo: ", c.Saldo)
	return c.Saldo
}

func main() {
	conta := Conta{
		Saldo: 100,
	}
	conta.emprestimo(200)
	println("Valor do saldo na variável conta: ", conta.Saldo)
}
