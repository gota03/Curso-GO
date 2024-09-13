package main

func main() {
	a := 10
	var ponteiro *int = &a
	*ponteiro = 20
	println(*ponteiro)

	// Um ponteiro é uma variável que armazena o endereço de memória de outra variável. Ao invés de armazenar um valor diretamente, o ponteiro armazena o local na memória onde o valor está guardado. Isso permite que modifique o valor original de uma variável dentro de funções ou métodos, ao invés de trabalhar com uma cópia.
	// & Permite acessar o endereço de memória da variável
	// * na declaração: Indica que estamos lidando com um ponteiro para um tipo. Exemplo: *int significa um ponteiro para uma variável inteira.
	// * ao acessar o valor: Desreferencia o ponteiro, ou seja, acessa o valor no endereço de memória para o qual o ponteiro aponta
}
