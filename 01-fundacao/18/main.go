// Pacote Principal
package main

// Importasndo pacotes Externos
import (
	"fmt"

	"github.com/devjuniorhanun/CursoGoLang/matematica"
	"github.com/google/uuid"
)

// Função Principal
func main() {
	// Chmando o Método Soma() no pacote matematica
	s := matematica.Soma(10, 20)
	// Instanciando a Classe Carro do pacote matematica
	carro := matematica.Carro{Marca: "Fiat"}

	// Imprimindo o método Andar() da Classe Carro
	fmt.Println(carro.Andar())
	// Imprimi o resultado do método soma
	fmt.Println("Resultado: ", s)
	// Chava a variável A do pacote matematica
	fmt.Println(matematica.A)
	// Imprimi um novo uuid
	fmt.Println(uuid.New())
}
