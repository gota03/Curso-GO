package main

import (
	"log"
	"net/http"
)

func main() {
	mux := http.NewServeMux()
	fileServer := http.FileServer(http.Dir("./Public"))
	fileServer2 := http.FileServer(http.Dir("./Public/Teste"))
	// o FileServer cria um manipulador HTTP que serve arquivos de um sistema de arquivos local. Ele é responsável por retornar arquivos para o cliente com base nas requisições.
	// O FileServer vai responder às requisições HTTP procurando arquivos no diretório
	// Está sendo configurado para servir arquivos do diretório ./Public. Isso significa que qualquer arquivo dentro desse diretório poderá ser acessado diretamente via URL.
	// http.Dir converte um diretório no sistema de arquivos (neste caso, o diretório "./Public") em um tipo http.FileSystem, que o FileServer pode utilizar.
	// O argumento "./Public" refere-se ao diretório onde estão os arquivos que serão servidos. O caminho "./" é relativo à localização do arquivo Go que você está executando. Se esse diretório não existir, o servidor não conseguirá servir arquivos corretamente.

	mux.Handle("/", fileServer)
	mux.Handle("/teste", fileServer2)
	// o diretório especificado será servido quando o servidor receber uma requisição para a raiz do servidor. Isso significa que qualquer arquivo dentro de ./Public pode ser acessado diretamente via URL.

	log.Fatal(http.ListenAndServe(":8080", mux))
	// A função log.Fatal é usada para registrar uma mensagem de erro e encerrar o programa se algo der errado.

	// log.Fatal():
	// Imprime uma mensagem de erro e encerra o programa imediatamente.
	// Não executa funções adiadas (defer).
	// Usa os.Exit(1) para finalizar o programa.

	// panic():
	// Lança um erro que interrompe o fluxo normal do programa e pode ser capturado com recover().
	// Executa funções adiadas (defer) antes de encerrar o programa.
	// Propaga o erro pela pilha de chamadas.
}
