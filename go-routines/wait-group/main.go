package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

func main() {

	wg.Add(3) //IMPORTANTE: Setear cuantas go-routinas vamos a controlar.

	go contar(25) // Se ejecuta como un nuevo sub-proceso.
	go contar(50) // Se ejecuta como un nuevo sub-proceso.
	go contar(70) // Se ejecuta como un nuevo sub-proceso.

	fmt.Println("\nfunc main() => Termino de ejecutarse")

	wg.Wait()
}

/*
	PROBLEMA: Si necesitamos conocer el resultado del proceso de la funcion contar()
	no vamos a poder, porque tanto la funcion contar() como main() se estan ejecutando como gorutinas. Pero main()
	se ejecuta como proceso princiapl y contar() como sub-proceso.

	SOLUCION: Para poder tener el control de lo que pasa
	en la funcion contar(), desde la funcion main(), debemos utilizar [WaitGroup].
*/
func contar(hasta int) {

	fmt.Printf("\nfunc contar(%v) => Comienza a ejecutarse\n", hasta)

	for i := 0; i <= hasta; i++ {
		fmt.Print(i)
	}

	fmt.Printf("\nfunc contar(%v) => Termino de ejecutarse\n", hasta)

	wg.Done() // Avisamos que esta funcion termino su proceso.
}
