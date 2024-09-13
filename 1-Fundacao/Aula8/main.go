package main

import (
	"errors"
	"fmt"
)

func main() {
	resultado, err := soma(4, 9)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(resultado)

	// para que a mensagem de erro seja exibida é preciso utilizar o pacote fmt
}

func soma(a, b int) (int, error) {
	if a+b > 10 {
		return 0, errors.New("Soma maior que 10")
	}
	return a + b, nil
}

// funções em go podem ter retorno com mais de um valor com tipos distintos
