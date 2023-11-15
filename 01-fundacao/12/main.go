// Pacote Principal
package main

// Função soma, que pega o apontamento da memoria
func soma(a, b *int) int {
	*a = 50
	*b = 50
	return *a + *b
}

// Função Principal
func main() {
	// Atribui 10 a minhaVar1
	minhaVar1 := 10
	// Atribui 20 a minhaVar2
	minhaVar2 := 20
	// Chama a função soma
	soma(&minhaVar1, &minhaVar2)

	println(minhaVar1)
	println(minhaVar2)
}
