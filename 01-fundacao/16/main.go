// Pacote Principal
package main

import "fmt"

// Função Principal
func main() {
	// Instanciando uma interface
	var minhaVar interface{} = "Meu Nome"
	// Imprimindo em string
	println(minhaVar.(string))
	// Tenta converter uma string em int
	res, ok := minhaVar.(int)
	// Imprimi o resultado da conversão
	fmt.Printf("O valor de res é %v e o resultado de ok é %v", res, ok)
	// Tenta converter uma string em int
	res2 := minhaVar.(int)
	// Imprimi o resultado da conversão, gerando um erro
	fmt.Printf("O valor de res2 é %v", res2)
}
