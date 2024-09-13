package main

import (
	"io"
	"net/http"
)

func main() {
	client := http.Client{}
	// Permite configurar opções como tempo limite, transporte (como proxies ou customizações de transporte), e outras opções.
	
	req, err := http.NewRequest("GET", "http://google.com", nil)
	// Parâmetros da função NewRequest:

	// method: O método HTTP para a requisição, como "GET", "POST", "PUT", etc.
	// url: O URL para onde a requisição deve ser enviada.
	// body: O corpo da requisição, que deve implementar a interface io.Reader. No caso de uma requisição GET, isso pode ser nil se não houver corpo.

	// Retorno da função NewRequest:

	// Retorna um ponteiro para um http.Request e um possível erro.

	if err != nil {
		panic(err)
	}
	req.Header.Set("Acept", "application/json")
	// O Header é um campo da struct Request que é do tipo Header que possui o método Set
	// O método Set é utilizado para definir um valor para um cabeçalho específico.
	// Configurando o cabeçalho Accept da requisição para indicar que o cliente aceita respostas no formato JSON.

	res, err := client.Do(req)
	// envia a requisição HTTP e retorna a resposta
	
	if err != nil {
		panic(err)
	}
	result, err := io.ReadAll(res.Body)
	if err != nil {
		panic(err)
	}
	defer res.Body.Close()
	println(string(result))
}
