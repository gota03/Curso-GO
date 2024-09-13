package main

import (
	"io"
	"net/http"
	"time"
)

func main() {
	client := http.Client{Timeout: time.Millisecond}
	// Timeout específica um tempo lime para requisições feitas por esse Client

	res, err := client.Get("http://google.com")
	// Caso a requisição não seja feita no tempo limite especificado, ocorrerá um erro

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
