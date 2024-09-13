package main

import (
	"net/http"
	"text/template"
)

type Curso struct {
	Nome         string
	CargaHoraria int
}

type ListaCursos []Curso

func main() {
	http.HandleFunc("/", Cursos)
	// Ao acessar a url "/" será chamada a função Cursos onde a função cria um template que renderiza as informações contidas no slice ListaCursos

	http.ListenAndServe(":8080", nil)
}

func Cursos(res http.ResponseWriter, req *http.Request) {
	temp := template.Must(template.New("template.html").ParseFiles("template.html"))
	err := temp.Execute(res, ListaCursos{
		{"GO", 20},
		{"Python", 15},
	})
	if err != nil {
		panic(err)
	}

}
