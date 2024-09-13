package main

import "fmt"

func main() {
	slice := []int{10, 20, 30, 40, 50}

	//  Um slice é uma referência a uma sequência de elementos que pode crescer ou diminuir dinamicamente. Um slice não tem um tamanho fixo, mas sim uma capacidade, que pode ser maior ou igual ao seu comprimento. Internamente, um slice contém um ponteiro para um array subjacente, um comprimento (length) e uma capacidade (capacity).

	// O array subjacente é onde os dados reais de um slice são armazenados.

	// O tamanho de um slice é o número de elementos que ele atualmente contém.

	// A capacidade de um slice é o número total de elementos que ele pode armazenar sem precisar realocar o array subjacente.

	// Se o número de elementos exceder a capacidade, o slice cria um novo array subjacente com o dobro da capacidade antiga (ou algo próximo disso), copia os elementos antigos para o novo array e adiciona o novo elemento.

	//Ao realizar o fatiamento de um slice declarando um ponto de início, a capacidade do novo slice será igual ao número de elementos restantes no array subjacente, a partir do índice de início até o final.

	fmt.Printf("len=%d cap=%d %v\n", len(slice), cap(slice), slice)

	fmt.Printf("len=%d cap=%d %v\n", len(slice[:0]), cap(slice[:0]), slice[:0])

	fmt.Printf("len=%d cap=%d %v\n", len(slice[:2]), cap(slice[:2]), slice[:2])

	fmt.Printf("len=%d cap=%d %v\n", len(slice[4:]), cap(slice[4:]), slice[4:])
}
