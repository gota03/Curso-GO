package main

import "fmt"

func main() {
	var array [3]int
	array[0] = 10
	array[1] = 20
	array[2] = 30

	//  Um array em Go tem um tamanho fixo, definido no momento da sua criação. Esse tamanho não pode ser alterado.

	// A memória para o array é alocada de forma contígua e diretamente no momento da criação do array.

	// Arrays são mais eficientes em termos de memória quando o tamanho é conhecido e fixo

	for i, v := range array {
		fmt.Printf("O valor do indice é %d e o valor é %d\n", i, v)
	}

	// o i recebe o índice do elemento atual do array (ou slice, mapa, etc.).
	// o v recebe o valor do elemento na posição do índice i.
	// O := declara essas variáveis e as inicializa com os valores apropriados para cada iteração do loop.
}
