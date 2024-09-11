package main

import (
	"encoding/json"
	//"encoding/json": Usado para converter (marshalling) dados entre JSON e structs, ou seja, faz o parsing do JSON recebido da API.

	"os"
	// "os": Pacote usado para interagir com o sistema operacional, aqui sendo usado para acessar argumentos de linha de comando e criar arquivos.
)

type Conta struct {
	Numero int `json:"n"`
	Saldo  int `json:"s"`
	// As tags são usadas para mapear os nomes dos campos da struct para os respectivos nomes no formato JSON e quando quando a struct for convertida para JSON.
}

func main() {
	conta := Conta{
		Numero: 1,
		Saldo:  1000,
	}
	res, erro := json.Marshal(conta)

	// A função Marshal converte a struct conta para uma sequência JSON. O resultado é retornado como um slice de bytes ([]byte).

	if erro != nil {
		panic(erro)
	}
	println(string(res))

	erro = json.NewEncoder(os.Stdout).Encode(conta)

	// É criado um codificador JSON que escreve diretamente para a saída padrão (os.Stdout). O método Encode converte a struct conta em JSON e a imprime imediatamente no console.

	if erro != nil {
		panic(erro)
	}

	jsonPuro := []byte(`{"Numero":2, "Saldo": 200}`)

	// É utilizado crases (``) para definir uma string literal de múltiplas linhas (raw string literal). Isso é útil para o JSON porque evita a necessidade de escapar caracteres especiais.

	var conta2 Conta
	err := json.Unmarshal(jsonPuro, &conta2)

	// &conta2 é passado por referência, pois o Unmarshal modifica a struct.

	if err != nil {
		panic(err)
	}
	println(conta2.Saldo)

	jsonModificado := []byte(`{"n": 3, "s": 300}`)
	var conta3 Conta
	err2 := json.Unmarshal(jsonModificado, &conta3)
	if err2 != nil {
		panic(err2)
	}
	println(conta3.Saldo)
}
