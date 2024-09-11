package main

const nome = "Mateus"

var (
	a bool
	b string
	c int
	d float64
)

func main() {
	//var e = "variavel local"
	// uma variável declarada em escopo local precisa ser utilizada

	f := "inferência de tipo"
	f = "mudando valor"

	// ao usar := o tipo da variável vai ser inferido de acordo com o valor atribuído, porém só pode ser utilizado ao declarar a variável, após isso deve ser utilizado o = para alterar o valor da variável
	// o := só pode ser utilizado em variáveis com escopo local

	println(f)
}
