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
	// Criando um array com 3 posições
	var meuArray [3]int
	meuArray[0] = 10 // Instanciando 10 no primeira posição do array
	meuArray[1] = 20 // Instanciando 20 no segunda posição do array
	meuArray[2] = 30 // Instanciando 30 no terceira posição do array

	// Percorrendo o array
	for i, v := range meuArray {
		// Imprimindo o dados do array
		fmt.Printf("O valor do indice é %d e o valor é %d\n", i, v)
	}
}
