package main

import (
	"fmt"
)

// Increase ...
func Increase(a int) {
	a++
	fmt.Println("Direccion de memoria de [a]:", &a)
}

// IncreaseWithPointers ... usando punteros
func IncreaseWithPointers(puntero *int) {
	*puntero++
	fmt.Println("Direccion de memoria de [b]:", &puntero)
}

func main() {
	valor := 19

	// Caso 1 - Sin utilizar punteros.
	Increase(valor) // intentamos modificamos el valor.
	fmt.Println("Resultado:", valor)

	// Caso 2 - Utilizando punteros.
	IncreaseWithPointers(&valor) // intentamos modificamos el valor aplicando punteros.
	fmt.Println("Resultado:", valor)

}
