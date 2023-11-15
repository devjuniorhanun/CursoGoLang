// Pacote Principal
package main

// Importação de Pacotes externos
import (
	"fmt"
)

// Criando um struct ( class ) de Endereço
type Endereco struct {
	Logradouro string
	Numero     int
	Cidade     string
	Estado     string
}

// Criando um struct ( class ) de Cliente
type Cliente struct {
	Nome  string
	Idade int
	Ativo bool
	Endereco
}

// Criando um Método para o struct de Cliente
func (c Cliente) Desativar() {
	// Atribuindo false para o argumento da struct
	c.Ativo = false
	// Imprimi
	fmt.Printf("O cliente %s foi desativado", c.Nome)
}

// Função Principal
func main() {
	// Instancia um novo cliente
	cliente := Cliente{
		Nome:  "Cliente",
		Idade: 30,
		Ativo: true,
	}
	// Muda o valor do argumento Ativo
	cliente.Ativo = false
	// Chama o método Desativar()
	cliente.Desativar()
}
