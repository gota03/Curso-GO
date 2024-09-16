package main

import (
	"context"
	"time"
)

func main() {
	ctx := context.Background()
	ctx, cancel := context.WithTimeout(ctx, time.Second*5)
	explodirBomba(ctx)
	defer cancel()
}

func explodirBomba(ctx context.Context) {
	select {
	// a select é usada para esperar por um ou mais canais. Ela bloqueia até que um dos canais esteja pronto para enviar ou receber um valor
	case <-ctx.Done():
		// O canal ctx.Done() é fechado quando o contexto é cancelado (o que pode ocorrer após o tempo especificado na função WithTimeout). Se esse canal estiver pronto antes do time.After, a mensagem "Bomba não explodiu" é impressa e a função retorna.

		println("Bomba não explodiu")
		return
	case <-time.After(time.Second * 7):
		// After espera que a duração decorra e então envia a hora atual no canal retornado
		// Cria um canal que será enviado um valor após 7 segundos. Se 7 segundos se passarem antes do contexto ser cancelado, a mensagem "Bomba explodiu" é impressa e a função retorna.
		
		println("Bomba explodiu")
		return
	}
}
