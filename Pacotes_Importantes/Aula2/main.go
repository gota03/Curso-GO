package main

import (
	"io"
	"net/http"
)

func main() {
	req, erro := http.Get("https://www.google.com")

	//A função http.Get retorna dois valores:
	//1 -> Um objeto *http.Response que contém os dados da resposta, como o status HTTP, cabeçalhos e o corpo da resposta (Body).
	//2 -> Caso ocorra um erro na requisição, ele é retornado nesta variável.

	if erro != nil {
		panic(erro)
	}
	defer req.Body.Close()

	// Defer garante que a linha de código seja executada por último

	res, erro := io.ReadAll(req.Body)

	// A função ReadAll lê todo o conteúdo do corpo da resposta HTTP (req.Body) e retorna os dados lidos como um slice de bytes ([]byte).

	if erro != nil {
		panic(erro)
	}

	// Fechar o Body após ler os dados, para liberar recursos do sistema.

	println(string(res))
}
