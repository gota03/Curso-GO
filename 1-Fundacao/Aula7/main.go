package main

func main() {
	salarios := map[string]int{"Mateus": 4000, "Lyandra": 5000, "Renato": 3000, "Jorge": 2000}

	// o map cria uma hash table que possui pares de chave e valor não ordenados, o retorno é um objeto do tipo map

	delete(salarios, "Jorge")

	// a função delete serve para deletar uma chave da hash table

	println(salarios["Jorge"])

	// ao acessar uma chave que não existe a saída é 0 ou ""

	salarios2 := make(map[string]float64)

	// a função make aloca e inicializa um objeto do tipo slice, map ou channel
	// ao criar um map a função make retorna um map vazio
	// na função make também pode ser passada a capacidade para o map e tamanho e capacidade para um slice

	print(salarios2)

	// for nome, salario := range salarios {
	// 	println(nome, salario)
	// }

	for _, salario := range salarios {
		print(salario)
	}
	// o blank identifier serve para ignorar algum valor no iterável percorrido
}
