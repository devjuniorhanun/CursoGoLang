// Pacote Principal
package main

// Importação de Pacotes externos
import (
	"fmt"
)

// Criando uma constante
const a = "Hello, World!"

// Criando um tipo inteiro
type ID int

// Criando Variáveis Globais
var (
	b bool    = true   // Variável booleana
	c int     = 10     // Variável inteira
	d string  = "Nome" // Variável string ( texto )
	e float64 = 1.2    // Variável float ( ponto flutuante )
	f ID      = 1      // Tipo Inteiro
)

// Função Principal
func main() {
	// Imprime o tido do dados
	fmt.Printf("O tipo de E é %T\n", f)
}
