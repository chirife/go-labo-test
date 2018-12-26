package main

import "fmt"

func main() {

	fmt.Println("funcion main() ejecutada")
}

// Se ejecuta antes del main.
func init() {
	fmt.Println("funcion init() ejecutada")
}
