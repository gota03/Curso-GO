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

// ListaCursos é um tipo de slice ([]Curso) onde cada elemento é do tipo Curso.

func main() {
	temp := template.Must(template.New("template.html").ParseFiles("template.html"))
	//  Quando usa ParseFiles, o nome do template que você especificou ao chamar template.New é usado para associar o arquivo carregado a um template específico.
	// O nome usado em template.New deve corresponder ao nome que você usa ao chamar template.ParseFiles porque ParseFiles associa o conteúdo dos arquivos aos templates que foram criados.

	err := temp.Execute(os.Stdout, ListaCursos{
		{"GO", 60},
		{"Python", 50},
		{"Javascript", 35},
		// {} é a sintaxe de inicialização de slices e structs em Go.
		// Cria uma instância do tipo ListaCursos e inicializa o slice com três elementos, onde cada elemento é uma instância da estrutura Curso.
	})
	if err != nil {
		panic(err)
	}
}
