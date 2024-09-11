package main

import (
	"encoding/json"
	// "encoding/json": Usado para converter (marshalling) dados entre JSON e structs, ou seja, faz o parsing do JSON recebido da API.

	"fmt"
	"io"

	//"io": Usado para manipulação de I/O (entrada/saída), neste caso para ler o corpo da resposta

	"net/http"
	// "net/http": Pacote que permite fazer requisições HTTP.

	"os"
	// "os": Pacote usado para interagir com o sistema operacional, aqui sendo usado para acessar argumentos de linha de comando e criar arquivos.
)

type ViaCep struct {
	Cep         string `json:"cep"`
	Logradouro  string `json:"logradouro"`
	Complemento string `json:"complemento"`
	Unidade     string `json:"unidade"`
	Bairro      string `json:"bairro"`
	Localidade  string `json:"localidade"`
	Uf          string `json:"uf"`
	Estado      string `json:"estado"`
	Regiao      string `json:"regiao"`
	Ibge        string `json:"ibge"`
	Gia         string `json:"gia"`
	Ddd         string `json:"ddd"`
	Siafi       string `json:"siafi"`
}

func main() {
	for _, url := range os.Args[1:] {
		// os.Args[1:]: Captura os argumentos passados via linha de comando. O primeiro argumento os.Args[0] é o nome do programa, então a partir de os.Args[1:] são considerados os parâmetros (neste caso, o CEP).

		req, err := http.Get("http://viacep.com.br/ws/" + url + "/json/")
		if err != nil {
			fmt.Fprintf(os.Stderr, "Erro ao fazer a requisição: %v\n", err)
		}
		defer req.Body.Close()
		res, err := io.ReadAll(req.Body)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Erro ao ler resposta: %v\n", res)
		}
		var cep ViaCep
		err = json.Unmarshal(res, &cep)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Erro ao fazer o unmarshal da resposta: %v\n", err)
			// A função fmt.Fprintf é usada para escrever uma string formatada em um destino específico, como um arquivo ou uma saída (por exemplo, os.Stderr ou os.Stdout).
			// O destino pode ser qualquer coisa que implemente a interface io.Writer, como arquivos, buffers, ou saídas padrão como os.Stdout (saída padrão) e os.Stderr (saída de erro padrão).
		}
		file, err := os.Create("cidade.txt")
		if err != nil {
			fmt.Printf("Erro ao criar arquivo: %v", err)
		}
		defer file.Close()
		_, err = file.WriteString(fmt.Sprintf("CEP: %v\nCidade: %v\nBairro: %v", cep.Cep, cep.Localidade, cep.Bairro))
		// fmt.Sprintg retorna a string formatada, permitindo que você a use posteriormente (por exemplo, atribuindo a uma variável, ou passando para outra função).

		if err != nil {
			fmt.Printf("Erro ao escrever no arquivo: %v", err)
		}
	}
}
