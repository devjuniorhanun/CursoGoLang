// Pacote Principal
package main

// Importação de Pacotes externos
import (
	"fmt"
)

// Função Principal
func main() {
	// Imprimi a função soma
	fmt.Println(sum(1, 3, 45, 6, 34, 654, 654, 7645, 534, 543, 543, 543))
}

// Função Sum()
// Recebe intinitos valores inteiros.
// Retorna o valor inteiro
func sum(numeros ...int) int {
	// Inicializa o contato
	total := 0
	// Percorre o array
	for _, numero := range numeros {
		// Realiza a soma do total + o número
		total += numero
	}
	// Retorna o total
	return total
}
