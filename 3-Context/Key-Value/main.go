package main

import (
	"context"
	"fmt"
)

func main() {
	ctx := context.Background()
	ctx = context.WithValue(ctx, "nome", "mateus")
	// context.WithValue permite que você armazene dados no contexto, como pares chave-valor, para que eles possam ser acessados posteriormente em outras funções.
	ReservaHotel(ctx)
}

func ReservaHotel(ctx context.Context) {
	nome := ctx.Value("nome")
	// Recupera o valor associado à chave "nome" do contexto passado. Se a chave existir no contexto, o valor correspondente é retornado (neste caso, "mateus").

	fmt.Println(nome)
}
