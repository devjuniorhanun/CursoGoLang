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

// Criando um interface de Pessoa
type Pessoa interface {
	// Método de Desativar()
	Desativar()
}

// Criando um struct ( class ) de Empresa
type Empresa struct {
	Nome string
}

// Método Desativar()
func (e Empresa) Desativar() {
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

func Desativacao(pessoa Pessoa) {
	pessoa.Desativar()
}

// Função Principal
func main() {
	// Instancia um novo cliente
	cliente := Cliente{
		Nome:  "Cliente",
		Idade: 41,
		Ativo: true,
	}
	fmt.Println(cliente)
	// Instanciando uma nova Empresa
	minhaEmpresa := Empresa{}

	// Desativa a Empresa
	Desativacao(minhaEmpresa)
	fmt.Println(minhaEmpresa)
}
