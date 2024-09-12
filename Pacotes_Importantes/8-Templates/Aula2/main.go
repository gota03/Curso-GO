package main

import (
	"os"
	"text/template"
)

type Curso struct {
	Nome         string
	CargaHoraria int
}

func main() {
	curso := Curso{
		Nome:         "GO",
		CargaHoraria: 60,
	}
	template := template.Must(template.New("curso-go").Parse("Curso: {{.Nome}} - Carga Horária: {{.CargaHoraria}}"))
	// É uma função que recebe um valor do tipo *template.Template e um erro (error). Se o erro for nil, ela retorna o template. Se o erro não for nil, ela faz o programa parar e imprime o erro.
	// Simplifica o tratamento de erros ao trabalhar com templates. É frequentemente usada para garantir que a criação e o parsing de templates sejam bem-sucedidos e para simplificar o código, especialmente em exemplos e scripts onde erros não são esperados.

	erro := template.Execute(os.Stdout, curso)
	if erro != nil {
		panic(erro)
	}

}
