package main

import (
	"context"
	"io"
	"net/http"
	"time"
)

func main() {
	ctx := context.Background()
	// Cria um novo contexto vazio que é a raiz da árvore de contextos. Este contexto é frequentemente usado como o contexto inicial para outras operações ou contextos derivados.

	ctx, cancel := context.WithTimeout(ctx, time.Millisecond)
	// Definindo um contexto com base no tempo
	// Cria um novo contexto derivado do contexto pai (parent) com um limite de tempo (timeout). Quando o tempo expira, o contexto é automaticamente cancelado.
	// cancle é uma função que pode ser chamada para cancelar o contexto manualmente antes do tempo expirar.

	defer cancel()
	// cancel() é chamada com defer para garantir que o contexto seja cancelado, liberando quaisquer recursos associados a ele.

	req, err := http.NewRequestWithContext(ctx, "GET", "http://google.com", nil)
	// Criando uma requisição com um contexto
	
	if err != nil {
		panic(err)
	}
	res, err := http.DefaultClient.Do(req)
	// É uma instância pré-configurada e global do tipo http.Client fornecida pela biblioteca padrão do Go.
	// Uma maneira rápida e fácil de fazer requisições HTTP sem a necessidade de criar e configurar um cliente HTTP personalizado.

	if err != nil {
		panic(err)
	}	
	defer res.Body.Close()
	result, err := io.ReadAll(res.Body)
	if err != nil {
		panic(err)
	}
	println(string(result))
}
