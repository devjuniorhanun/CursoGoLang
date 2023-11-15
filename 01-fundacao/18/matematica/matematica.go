// Pacote matematica
package matematica

// Método Soma().
// Pode receber interiros ou ponto flutuante
func Soma[T int | float64](a, b T) T {
	// Retorna a soma dos valores
	return a + b
}

// Iniializa uma variável inteira
var A int = 10

// Cria uma Classe Carro
type Carro struct {
	Marca string
}

// Cria um método para a classa Carro
func (c Carro) Andar() string {
	// Retorna o que o carro fez
	return "Carro andando"
}
