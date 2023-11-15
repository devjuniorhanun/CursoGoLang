// Pacote Principal
package main

// Importação de Pacotes externos
import (
	"fmt"
)

// Função Principal
func main() {
	// Criando um Maps (Array) com chave e valor
	salarios := map[string]int{"Júnior": 1000, "João": 2000, "Maria": 3000}
	// Excluir um index do maps
	delete(salarios, "Júnior")
	// Adiciona um novo index
	salarios["Jún"] = 5000

	// sal := make(map[string]int)
	// sal1 := map[string]int{}
	// sal1["Júnior"] = 1000

	// Percorre o map, mostrando a chave e o valor
	for nome, salario := range salarios {
		fmt.Printf("O salario de %s é %d\n", nome, salario)
	}

	// Percorre o map, ocultando a chave ou valor
	for _, salario := range salarios {
		fmt.Printf("O salario é %d\n", salario)
	}
}
