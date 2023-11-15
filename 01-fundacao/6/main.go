// Pacote Principal
package main

// Importação de Pacotes externos
import (
	"fmt"
)

// Função Principal
func main() {
	// Criando um Slices (Array) infinito
	s := []int{10, 20, 30, 50, 60, 70, 80, 90, 100}
	// Imprimi o slice informando to tamanho do slice, e a capacidade que ele tem
	fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)

	fmt.Printf("len=%d cap=%d %v\n", len(s[:0]), cap(s[:0]), s[:0])

	fmt.Printf("len=%d cap=%d %v\n", len(s[:4]), cap(s[:4]), s[:4])

	fmt.Printf("len=%d cap=%d %v\n", len(s[2:]), cap(s[2:]), s[2:])

	// Adiciona um novo item ao slice
	s = append(s, 110)
	fmt.Printf("len=%d cap=%d %v\n", len(s[:2]), cap(s[:2]), s[:2])
}
