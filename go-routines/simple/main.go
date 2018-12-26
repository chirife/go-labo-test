package main

import (
	"fmt"
)

func main() {

	go contar(200) // Se ejecuta a parte como un sub-proceso.

	fmt.Println("\nFin del proceso principal")

}

/*
	PROBLEMA: Si necesitamos conocer el resultado del proceso de la funcion contar()
	no vamos a poder, porque tanto la funcion contar() como main() se estan ejecutando como gorutinas. Pero main()
	se ejecuta como proceso princiapl y contar() como sub-proceso.

	SOLUCION: Para poder tener el control de lo que pasa
	en la funcion contar() debemos utilizar [WaitGroup].
*/
func contar(hasta int) {

	for i := 0; i < hasta; i++ {
		fmt.Printf("contando (%v)", i)
	}

	fmt.Println("\nTerminamos de contar hasta ", hasta)
}
