package main

func soma(a, b *int) int {
	// Os parâmetros a e b são ponteiros para inteiros (*int), ou seja, eles armazenam endereços de memória, não valores diretos.
	// * Usado para acessar o valor armazenado no endereço de memória apontado por um ponteiro.
	*a = 40
	*b = 60
	return *a + *b
}

func main() {
	a := 10
	b := 20
	println(soma(&a, &b))
	// Ao usar &a, &b isso significa que a função não recebe uma cópia dos valores, mas sim a referência (endereço) onde os valores de a e b estão armazenados.
	println(a)
	println(b)
}
