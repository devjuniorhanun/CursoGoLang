// Pacote Principal
package main

// Criando um tipo de Int
type MyNumber int

// Criando um interface de números interiros ou com ponto flutuantes
type Number interface {
	~int | ~float64
}

// Criando um método de numeros, e um array array (Chave String)  (Valor Númerico)
func Soma[T Number](m map[string]T) T {
	// Criando uma variável de numerica
	var soma T
	// Percorre o mapa de string
	for _, v := range m {
		// Incrementa a variável soma
		soma += v
	}
	// Retorna a soma
	return soma
}

// Cria um método de comparação
func Compara[T comparable](a T, b T) bool {
	// Verifica se os dois dados, são iguaks
	if a == b {
		// Se for retorna verdadeiro
		return true
	}
	// Se não retorna falso
	return false
}

// Função Principal
func main() {
	// Instancia o primeiro array (Chave String)  (Valor Númerico)
	m := map[string]int{"Júnior": 1000, "João": 2000, "Maria": 3000}
	// Instancia o primeiro array (Chave String)  (Valor Númerico)
	m2 := map[string]float64{"Júnior": 100.20, "João": 2000.3, "Maria": 300.0}
	// Instancia o primeiro array (Chave String)  (Valor Númerico)
	m3 := map[string]MyNumber{"Júnior": 1000, "João": 2000, "Maria": 3000}
	// Imprimi o resutlado da soma do primeiro array
	println(Soma(m))
	// Imprimi o resutlado da soma do segundo array
	println(Soma(m2))
	// Imprimi o resutlado da soma do terceiro array
	println(Soma(m3))
	// Imprimi a compração entre dois números
	println(Compara(10, 10))
}
