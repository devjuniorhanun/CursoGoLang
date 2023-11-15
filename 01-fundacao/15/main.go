// Pacote Principal
package main

import "fmt"

// Função Principal
func main() {
	// Atribui 10 a interface de x
	var x interface{} = 10
	// Atribui Hello, World! a interface de y
	var y interface{} = "Hello, World!"
	// Passando x ao método
	showType(x)
	// Passando y ao método
	showType(y)
}

// Criando um método de interface vazia
func showType(t interface{}) {
	// Imprimi o tipo do dado
	fmt.Printf("O tipo da variavel é %T e o valor é %v\n", t, t)
}
