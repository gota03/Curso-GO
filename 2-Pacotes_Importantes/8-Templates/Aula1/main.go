package main

import (
	"os"
	"text/template"
	// Pacote que fornece funções para trabalhar com templates de texto. Ele permite criar templates que podem ser preenchidos com dados dinâmicos.
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
	tmp := template.New("Curso-Template")
	// cria um novo template.
	// um template é um mecanismo para gerar texto dinâmico a partir de um modelo (template) e dados fornecidos. Os templates permitem que você crie conteúdo textual (como HTML, e-mails, relatórios, etc.) de maneira flexível e reutilizável.

	tmp, _ = tmp.Parse("Curso: {{.Nome}} - Carga Horária: {{.CargaHoraria}}")
	// Define o conteúdo do template. No template, {{.Nome}} e {{.CargaHoraria}} são placeholders que serão substituídos pelos valores correspondentes da instância da estrutura Curso.

	err := tmp.Execute(os.Stdout, curso)
	// Executa o template, substituindo os placeholders pelos valores da estrutura curso.
	// os.Stdout é o destino para a saída do template, que neste caso é o console.
	if err != nil {
		panic(err)
	}
}
