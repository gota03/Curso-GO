package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	// Criando arquivo
	arquivo, erro := os.Create("arquivo.txt")
	if erro != nil {
		panic(erro)
		// A função panic(erro) é chamada, interrompendo a execução do programa e exibindo o erro.
	}
	println("Arquivo criado com sucesso")

	// Escrevendo no arquivo
	tamanho, erro := arquivo.WriteString("Preenchendo o arquivo")
	if erro != nil {
		panic(erro)
	}
	fmt.Printf("Tamanho do arquivo: %d bytes", tamanho)

	// Lendo arquivo
	conteudo_arquivo, erro := os.ReadFile("arquivo.txt")
	if erro != nil {
		panic(erro)
	}
	fmt.Printf("\nContéudo do arquivo: %v\n", string(conteudo_arquivo))
	arquivo.Close()

	// Lendo arquivo de byte em byte
	file, erro := os.Open("arquivo.txt")

	// A função Opne abre o arquivo nomeado para leitura. Se for bem-sucedido, os métodos no arquivo retornado podem ser usados para leitura;

	if erro != nil {
		panic(erro)
	}
	reader := bufio.NewReader(file)

	// A função NewReader é usada para criar um novo objeto Reader bufferizado. Este NewReader usa um buffer interno para ler dados de forma mais eficiente de uma fonte de entrada, como arquivos ou conexões de rede.
	// O NewReader cria um objeto que encapsula a fonte de dados (como um arquivo) e adiciona um buffer interno.
	// Cria um leitor bufferizado que melhora a eficiência ao ler dados de forma mais rápida e em blocos.

	buffer := make([]byte, 10)

	// Um buffer é essencialmente um espaço reservado na memória (normalmente na forma de um array ou slice) onde os dados são armazenados temporariamente antes de serem processados.
	// No caso de leitura de arquivos, em vez de ler cada caractere ou byte individualmente do arquivo, o sistema lê um bloco de bytes de uma só vez e armazena no buffer. Depois, os dados do buffer são processados aos poucos.

	for {
		i, erro := reader.Read(buffer)

		// A função Read é usada para ler dados do buffer criado pelo bufio.Reader ou de qualquer outra fonte que implemente a interface io.Reader. Ela lê uma quantidade de bytes de uma vez, até o tamanho do buffer que você especificar.
		// O parâmetro tem o papel duplo de local de armazenamento dos dados lidos e de controle da quantidade de dados a serem lidos por vez.
		// Retorna o número de bytes efetivamente lidos e, se o final do arquivo for alcançado, o erro retornado será io.EOF.
		// Lê uma parte dos dados do buffer e armazena no slice de bytes para ser processado.

		if erro != nil {
			break
		}
		fmt.Println(string(buffer[:i]))
		// Garante que apenas os bytes lidos (até i) sejam convertidos em string e exibidos, ignorando qualquer dado extra no buffer que não tenha sido preenchido.
	}
	erro = os.Remove("arquivo.txt")
	if erro != nil {
		panic(erro)
	}
}
