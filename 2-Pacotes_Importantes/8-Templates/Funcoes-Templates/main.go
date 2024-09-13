package main

import (
	"os"
	"strings"
	"text/template"
)

type Curso struct {
	Nome         string
	CargaHoraria int
}

type ListaCursos []Curso

func main() {
	templates := []string{
		"header.html",
		"content.html",
		"footer.html",
	}
	temp := template.New("content.html")
	temp.Funcs(template.FuncMap{"ToUpper": strings.ToUpper})
	// a função Funcs é usada para adicionar uma mapa de funções personalizadas que podem ser usadas nos templates.
	// O template.FuncMap associa uma string (neste caso "ToUpper") a uma função (neste caso strings.ToUpper).
	// Ao fazer isso, você está disponibilizando a função ToUpper para uso dentro dos templates, permitindo que qualquer string no template seja convertida para maiúsculas.

	temp = template.Must(temp.ParseFiles(templates...))

	err := temp.Execute(os.Stdout, ListaCursos{
		{"GO", 20},
		{"Python", 15},
	})
	if err != nil {
		panic(err)
	}
}
