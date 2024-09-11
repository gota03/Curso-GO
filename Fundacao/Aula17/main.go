package main

import "fmt"

type NumberTest int

type Number interface {
	~int | ~float64

	//~ É usado para criar uma restrição de tipo que permite qualquer tipo que seja "similar" a um tipo específico, em vez de exigir uma correspondência exata. É uma forma de especificar que um tipo deve ser do mesmo tipo base ou ter um tipo que é baseado em um tipo específico.
}

// Generics permitem definir restrições de tipo para limitar os tipos que podem ser usados com um parâmetro de tipo genérico. Isso é feito usando interfaces.

func Soma[T Number](mapa map[string]T) T {
	// Generics permitem definir funções e tipos que podem operar em diferentes tipos de dados de maneira segura e reutilizável

	var soma T
	for _, valor := range mapa {
		soma += valor
	}
	return soma
}

func main() {
	mapa1 := map[string]int{"Mateus": 100, "Lyandra": 200}
	mapa2 := map[string]float64{"Renato": 300.0, "Jorge": 400.0}
	mapa3 := map[string]NumberTest{"Renato": 300.0, "Jorge": 400.0}
	fmt.Printf("Soma: %v", Soma(mapa1))
	fmt.Printf("Soma: %v", Soma(mapa2))
	fmt.Printf("Soma: %v", Soma(mapa3))
}
