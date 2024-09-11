package main

import "fmt"

func main() {
	var x interface{} = 10
	var y interface{} = "mateus"
	mostrarTipo(x)
	mostrarTipo(y)
	// Interfaces vazias podem trabalhar com qualquer tipo
}

func mostrarTipo(t interface{}) {
	fmt.Printf("O tipo é %T\nO valor é %v\n", t, t)
}
