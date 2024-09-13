package main

func main() {
	resultado := soma(10, 20, 30, 40, 50, 60)
	println(resultado)

	total := func() int {
		return soma(10, 20, 30, 40, 50) * 2
	}()
	println(total)

	//criando uma função anônima que devolve um inteiro e retorna o retorno da função soma multiplicado por 2
}

func soma(numeros ...int) int {
	total := 0

	for _, valor := range numeros {
		total += valor
	}
	return total
}

// as reticiências servem para dizer que o parâmetro da função pode receber muitos valores
// a variável se torna um iterável
