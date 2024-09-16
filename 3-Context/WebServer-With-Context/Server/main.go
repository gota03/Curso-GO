package main

import (
	"log"
	// Utilizado para registrar mensagens no console (log) durante o processamento da requisição.

	"net/http"
	"time"
)

func main() {
	http.HandleFunc("/", ResponseWithContext)
	http.ListenAndServe(":8080", nil)
}

func ResponseWithContext(res http.ResponseWriter, req *http.Request) {
	ctx := req.Context()
	// Obtém o contexto da requisição. O contexto gerencia prazos e cancelamentos, o que permite saber se a requisição foi encerrada ou cancelada.

	log.Println("Requisição iniciada")
	defer log.Println("Requisição encerrada")
	select {
	case <-time.After(time.Second * 3):
		// Retorna um canal que "desperta" após N segundos. Se nada interferir nesses segundos, o código executa essa opção.

		log.Println("Requisição processada com sucesso")
		res.Write([]byte("Resposta enviada"))
	case <-ctx.Done():
		//  é acionado quando o cliente cancela a requisição (por exemplo, fecha o navegador antes de receber a resposta). Se isso acontecer antes do tempo definido em segundos, a mensagem "Requisição cancelada pelo cliente" é registrada no log, e nenhuma resposta será enviada.

		log.Println("Requisição cancelada pelo cliente")
	}
}
