// Pacote Principal
package main

// Importação de Pacotes externos

// Função Principal
func main() {
	// Atribuindo 10 ao a
	// Memória -> Endreço -> Valor
	a := 10
	// Criando um ponteiro de inteiro de a
	var ponteiro *int = &a
	// Atribuindo 20 ao ponteiro e a
	*ponteiro = 20
	b := &a
	*b = 30
	println(a)
}
