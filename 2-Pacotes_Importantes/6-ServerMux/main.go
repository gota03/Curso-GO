package main

import "net/http"

type Blog struct {
	Title string
}

func main() {
	mux := http.NewServeMux()
	// ServeMux: É um roteador que permite associar rotas (URLs) a manipuladores específicos. O ServeMux gerencia qual função deve ser chamada para cada URL.
	// O ServeMux (abreviação de Serve Multiplexer) é um roteador de requisições HTTP. Ele é responsável por mapear URLs para os respectivos manipuladores (handlers). Em outras palavras, ele recebe as requisições e as direciona para a função ou struct que deve processá-las.
	// Como funciona: Quando você cria um ServeMux, ele age como uma tabela de roteamento. Você pode registrar diferentes caminhos (URLs) e associar cada um a uma função específica ou a um manipulador que atenda a essa requisição.
	// Por que usar?: O ServeMux é útil para organizar e gerenciar as rotas de um servidor web, especialmente em projetos maiores. Ele facilita a separação de responsabilidades e a criação de rotas dinâmicas.

	mux.HandleFunc("/", HomeHandler)
	// Aceita diretamente uma função que implementa a lógica de um manipulador de requisições (chamado de handler) que tem a assinatura padrão func(res http.ResponseWriter, req *http.Request).
	// É usado quando você deseja passar diretamente uma função como manipulador.

	mux.Handle("/blog", Blog{Title: "meu blog"})
	// Espera um handler que implementa a interface http.Handler, o que significa que o handler deve possuir o método ServeHTTP(res http.ResponseWriter, req *http.Request).
	// Em vez de uma função simples, você pode passar uma struct que implemente o método ServeHTTP. Isso é útil quando você deseja que a lógica do handler seja associada a uma instância de uma estrutura.
	// É usado quando você deseja passar um objeto que implementa a interface http.Handler, ou seja, que tenha o método ServeHTTP.

	// É preciso registrar todas as rotas antes de iniciar o servidor

	erro := http.ListenAndServe(":8080", mux)
	// Está dizendo explicitamente que o ServeMux será responsável por gerenciar as rotas do seu servidor.
	// Vai lidar com o roteamento, permitindo que registre rotas personalizadas com HandleFunc() ou Handle().
	// O servidor HTTP vai delegar cada requisição ao mux, que verifica a URL e chama o manipulador (handler) apropriado.

	if erro != nil {
		panic(erro)
	}

}

func HomeHandler(res http.ResponseWriter, req *http.Request) {
	// Um handler é qualquer função ou objeto que trata uma requisição HTTP e envia uma resposta de volta ao cliente.
	// Um handler é qualquer função ou tipo que pode lidar com uma requisição HTTP e enviar uma resposta. Ele pode ser uma função (com HandleFunc) ou uma struct que implemente ServeHTTP (com Handle).

	_, erro := res.Write([]byte("Testando mux"))
	if erro != nil {
		panic(erro)
	}
}

func (b Blog) ServeHTTP(res http.ResponseWriter, req *http.Request) {
	// Funções ou structs que implementam o método ServeHTTP(res http.ResponseWriter, req *http.Request). Eles são responsáveis por processar as requisições e retornar respostas.
	// O método ServeHTTP(res http.ResponseWriter, req *http.Request) faz parte da interface http.Handler no pacote net/http. Qualquer tipo que implemente esse método pode ser usado como um handler para lidar com requisições HTTP e pode ser registrada com mux.Handle()

	_, erro := res.Write([]byte(b.Title))
	if erro != nil {
		panic(erro)
	}
}
