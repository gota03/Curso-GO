package main

import (
	"bytes"
	"io"
	"net/http"
	"os"
	"time"
)

func main() {
	client := http.Client{Timeout: time.Second}
	jsonBuffer := bytes.NewBuffer([]byte(`"nome": "mateus"`))
	// NewBuffer retorna um ponteiro para o tipo Buffer do pacote bytes
	// O tipo Buffer por sua vez possui um metódo que implementa a interface io.Reader

	res, err := client.Post("http://google.com", "application/json", jsonBuffer)
	// Argumentos da função Post:
	// URL: "http://google.com", o destino da requisição.
	// Content-Type: "application/json", que especifica o tipo de conteúdo enviado (um JSON neste caso).
	// Body: jsonBuffer, que contém o corpo da requisição, criado anteriormente pelo bytes.NewBuffer.
	// Por o tipo Buffer possuir um método que implementa a interface io.Reader, ele satisfaz a condição e pode ser passado como argumento para o 3° parâmetro da função Post

	if err != nil {
		panic(err)
	}
	defer res.Body.Close()
	io.CopyBuffer(os.Stdout, res.Body, nil)
	// A função io.CopyBuffer copia dados de um leitor (io.Reader) para um gravador (io.Writer), utilizando um buffer intermediário. Neste caso:
	// Leitor (io.Reader): res.Body, que contém o corpo da resposta HTTP retornada pelo servidor.
	// Gravador (io.Writer): os.Stdout, que representa a saída padrão do sistema, geralmente o terminal.
	// O terceiro argumento é o buffer, que neste caso está definido como nil. Quando passado como nil, a função cria e gerencia seu próprio buffer internamente.
}
