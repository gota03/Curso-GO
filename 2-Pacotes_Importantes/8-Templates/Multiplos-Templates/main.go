package main

import (
	"os"
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
	temp := template.Must(template.New("content.html").ParseFiles(templates...))
	// O nome passado na função New será usado como referência para o template principal, ou seja, aquele que você deseja renderizar diretamente.
	// Todos os templates definidos nos arquivos passados para ParseFiles são parseados e armazenados no conjunto de templates.
	// O nome passado para New ("content.html") deve corresponder ao nome de um dos templates dentro dos arquivos listados em ParseFiles, pois este será o template principal que será renderizado quando você chamar a função Execute.
	err := temp.Execute(os.Stdout, ListaCursos{
		{"GO", 20},
		{"Python", 15},
	})
	if err != nil {
		panic(err)
	}
}
