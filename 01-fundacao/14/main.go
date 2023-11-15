// Pacote Principal
package main

// Criando um Classe struct Conta
type Conta struct {
	saldo int
}

// Criando um método de nova conta
func NewConta() *Conta {
	// Retorna o Saldo da conta
	return &Conta{saldo: 0}
}

// Criando um método de simular
func (c *Conta) simular(valor int) int {
	// Soma o valor ao saldo inicializado
	c.saldo += valor
	// Imprimi o saldo
	println(c.saldo)
	// Retorna o Saldo
	return c.saldo
}

// Função Principal
func main() {
	// Instancia uma nova conta, passando 100 ao saldo
	conta := Conta{saldo: 100}
	// Chama o método simular e passa 200
	conta.simular(200)
	// Imprimi o saldo
	println(conta.saldo)
}
