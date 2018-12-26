package main

import (
	"fmt"
)

// structs comunes
/****************************************************/
type nombre string
type edad int

// Metodo del struct nombre.
func (e nombre) mostrar() {
	fmt.Printf("Mi nombre es '%s'", e)
}

// Interface
/****************************************************/

type persona interface {
	saludar()
}

func interactuar(p persona) {
	p.saludar()

}

// Clase Paciente
/****************************************************/
type cliente struct {
	nombre nombre
	edad   edad
}

func (c cliente) saludar() {
	fmt.Printf("\nhola soy %v y tengo %v años.", c.nombre, c.edad)
}

// Clase Proveedor
/****************************************************/
type proveedor struct {
	nombre nombre
	edad   edad
}

func (p proveedor) saludar() {
	fmt.Printf("\nhola soy %v y tengo %v años.", p.nombre, p.edad)
}

/****************************************************/

func main() {

	c1 := cliente{"Marcelo", 33}
	p1 := proveedor{"Juan", 62}

	// Accediendo a los metodos de cada clase.
	// c1.saludar()
	// p1.saludar()

	//Utilizando el metodo de la interfaz.
	interactuar(c1)
	interactuar(p1)

	c1.nombre.mostrar()

}
