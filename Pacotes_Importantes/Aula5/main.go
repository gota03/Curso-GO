package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type ViaCep struct {
	Cep         string `json:"cep"`
	Logradouro  string `json:"logradouro"`
	Complemento string `json:"complemento"`
	Unidade     string `json:"unidade"`
	Bairro      string `json:"bairro"`
	Localidade  string `json:"localidade"`
	Uf          string `json:"uf"`
	Estado      string `json:"estado"`
	Regiao      string `json:"regiao"`
	Ibge        string `json:"ibge"`
	Gia         string `json:"gia"`
	Ddd         string `json:"ddd"`
	Siafi       string `json:"siafi"`
}

func main() {
	http.HandleFunc("/", BuscaCepHandler)
	//Registra um manipulador (handler) para uma rota HTTP específica. Ele mapeia o caminho "/" (a raiz do site) para a função BuscaCep.
	//O primeiro argumento ("/") é o caminho da URL. Neste caso, a raiz.
	//O segundo argumento (BuscaCep) é a função que será chamada quando alguém acessar esse caminho.

	erro := http.ListenAndServe(":8080", nil)
	//Inicia o servidor HTTP na porta especificada. Ele "escuta" por requisições HTTP.
	//O primeiro argumento (":8080") define a porta em que o servidor estará rodando (neste caso, 8080).
	//O segundo argumento (nil) representa o handler padrão. Como estamos usando http.HandleFunc para definir o manipulador da rota, o nil indica que o pacote http deve usar os manipuladores registrados previamente.

	if erro != nil {
		fmt.Println("Erro ao iniciar o servidor:", erro)
		panic(erro)
	}
}

func BuscaCepHandler(res http.ResponseWriter, req *http.Request) {
	//res http.ResponseWriter: Representa a resposta HTTP que será enviada ao cliente. Com ele, podemos enviar dados de volta para o navegador que fez a requisição.
	// Características:
	// Escrever Respostas: Permite escrever dados de resposta para o cliente, como o corpo da resposta e cabeçalhos HTTP.
	// Definir Status Code: Permite definir o código de status HTTP da resposta (por exemplo, 200 OK, 404 Not Found, 500 Internal Server Error).
	// Configurar Cabeçalhos: Permite definir cabeçalhos HTTP na resposta (por exemplo, Content-Type, Content-Length).

	//req *http.Request: Representa a requisição HTTP feita pelo cliente. Contém informações como o método (GET, POST), cabeçalhos, etc.
	// Características:
	// Ler Dados da Requisição: Contém informações sobre a requisição, como cabeçalhos, parâmetros de consulta (query parameters), corpo da requisição, e URL.
	// Informações de Requisição: Fornece detalhes sobre a requisição, como o método HTTP (GET, POST), o caminho da URL, e os dados enviados pelo cliente.

	//O asterisco (*) no parâmetro req *http.Request indica que o parâmetro é um ponteiro para uma instância da struct http.Request. Permite que o código acesse os dados da requisição de forma eficiente e sem cópias desnecessárias.

	if req.URL.Path != "/" {
		// req.URL.Path: Acessa o caminho da URL da requisição

		res.WriteHeader(http.StatusNotFound)
		// res.WriteHeader: Define o código de status da resposta HTTP. O código http.StatusNotFound (404) é usado para indicar que o recurso solicitado não foi encontrado.

		return
	}
	cepParam := req.URL.Query().Get("cep")
	// req.URL.Query(): Acessa os parâmetros de consulta (query parameters) da URL. Estes são os parâmetros que vêm depois do ? na URL

	if cepParam == "" {
		res.WriteHeader(http.StatusBadRequest)
		// Responde com o código de status 400 Bad Request (http.StatusBadRequest). Esse código é usado quando o servidor detecta que a requisição está malformada ou faltam informações necessárias

		return
	}

	cep, err := BuscaCep(cepParam)
	if err != nil {
		res.WriteHeader(http.StatusInternalServerError)
		//http.StatusInternalServerError (valor 500), que indica um erro interno do servidor.

		return
	}

	res.Header().Set("Content-Type", "application/json")
	// Content-Type: Informa o tipo de conteúdo do corpo da resposta, como JSON, HTML, ou XML.
	// Define o cabeçalho Content-Type da resposta HTTP para "application/json". O cabeçalho Content-Type informa ao cliente qual o formato dos dados que estão sendo retornados. Aqui, estamos informando que o corpo da resposta será um JSON.

	res.WriteHeader(http.StatusOK)
	// Define o código de status da resposta HTTP como 200 OK (http.StatusOK), indicando que a requisição foi processada com sucesso.

	// res.Write([]byte("Ola mundo"))
	//res.Write: Escreve a resposta que será enviada ao cliente. Neste caso, estamos escrevendo a string "Ola mundo" convertida em um array de bytes ([]byte).
	//O valor de retorno da função Write é o número de bytes escritos e um erro (se houver). Aqui, o número de bytes escritos não está sendo usado, então apenas o erro é capturado.

	erro := json.NewEncoder(res).Encode(cep)
	// NewEncoder: Cria um novo encoder JSON.
	// Esse encoder é configurado para escrever dados JSON diretamente no http.ResponseWriter (res), que é a saída da resposta HTTP. Permite codificar (converter) dados em formato JSON e escrevê-los diretamente no corpo da resposta HTTP, sem a necessidade de criar uma string JSON separada e depois escrever no res.
	// Encode: O método Encode recebe o valor de cep e o converte em JSON.
	// Depois de converter os dados em JSON, ele escreve esse JSON diretamente no http.ResponseWriter (res).

	if erro != nil {
		res.WriteHeader(http.StatusInternalServerError)
		// http.StatusInternalServerError (valor 500), que indica um erro interno do servidor.

		return
	}
}

func BuscaCep(cep string) (*ViaCep, error) {
	// * na declaração: Indica que estamos lidando com um ponteiro para um tipo. Exemplo: *ViaCep significa um ponteiro para uma struct ViaCep.

	req, erro := http.Get("http://viacep.com.br/ws/" + cep + "/json/")
	if erro != nil {
		fmt.Printf("Erro: %v", erro)
	}
	defer req.Body.Close()

	res, erro := io.ReadAll(req.Body)
	if erro != nil {
		fmt.Printf("Erro: %v", erro)
	}

	var c ViaCep
	erro = json.Unmarshal(res, &c)
	if erro != nil {
		fmt.Printf("Erro: %v", erro)
	}
	return &c, nil
	// Ao retornar o endereço de memória (&c), a função evita a sobrecarga de alocação de uma nova cópia da struct, o que melhora o desempenho
	// Retornar um ponteiro permite que a variável original possa ser modificada diretamente pelo código que receber o retorno
	// Se retornasse o valor da variável (*c), o Go teria que copiar toda a struct e o receptor teria apenas uma cópia e qualquer alteração feita na struct retornada não afetaria o original.
	// O endereço retornado por &c é o mesmo endereço onde os dados da struct ViaCep começam na memória.
	// Quando retorna &c, o que está retornando é o endereço de memória da instância de ViaCep armazenada em c. Esse endereço é o ponto de partida para acessar os dados dessa struct.
}
