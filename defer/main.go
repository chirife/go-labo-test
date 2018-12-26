package main

import (
	"fmt"
)

/* Defer */
func main() {

	/*
		[Defer]: Se usa para asegurar que una función es llamada
		posteriormente durante la ejecución del programa, generalmente
		con propósitos de limpieza. defer se usa regularmente donde en
		otros lenguajes se utilizaría ensure y finally
	*/

	uno := 20
	dos := 10

	defer suma(uno, dos)  // Se ejecuta tercero. (30)
	defer resta(uno, dos) // Se ejecuta segundo. (10)
	multi(uno, dos)       // Se ejecuta primero. (200)

	fmt.Println("final")

}

/* Suma */
func suma(x, y int) {
	fmt.Println(x + y)
}

/* Resta */
func resta(x, y int) {
	fmt.Println(x - y)
}

/* Multiplicación */
func multi(x, y int) {
	fmt.Println(x * y)
}
