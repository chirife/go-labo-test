package main

import "fmt"

func main() {

	puntos := []int{30, 20, 10} // Datos.
	suma(puntos...)             // parametros multiples.
	resta(puntos[0], puntos[1]) // parametros fijos.

}

/* Funciones con parametros multiples.*/
func suma(numeros ...int) {

	var contador int

	for _, v := range numeros {
		contador += v
	}

	fmt.Println(contador)
}

/* Funciones con parametros fijos.*/
func resta(num1 int, num2 int) {
	fmt.Println(num1 - num2)
}
