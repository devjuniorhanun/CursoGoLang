// Pacote Principal
package main

// Importação de Pacotes externos
import (
	"errors"
	"fmt"
)

// Função Principal
func main() {
	// Instancia a função soma
	valor, err := sum(50, 10)
	// Verifica se existe erro
	if err != nil {
		// Imprimi o erro
		fmt.Println(err)
	}
	// Imprime o resultado
	fmt.Println(valor)
}

// Função Sum()
// Recebe dois valores inteiros.
// Retorna o valor inteiro, e se caso tiver erro exibe o erro
func sum(a, b int) (int, error) {
	// Verifica a soma de a + b.
	if a+b >= 50 {
		// Ser for maior ou igual a 50, exibe a mensagem
		return 0, errors.New("A soma é maior que 50")
	}
	// Retorna a soma de a + b, e passa o erro como vazio
	return a + b, nil
}
