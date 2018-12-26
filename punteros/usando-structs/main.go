package main

import (
	"fmt"
)

type persona struct {
	nombre string
	edad   int
}

// Metodo que recupera el nombre de una persona.
func (p persona) getNombre() {
	fmt.Printf("%v tiene %v años.", p.nombre, p.edad)
}

// Metodo que modifica el nombre de una persona.
// NOTA: Para poder modificar algun dato del struct, es importante
// crear un puntero a la direccion de memoria del struct persona. Ej: (p *persona)
func (p *persona) setNombre(nuevoNombre string) {
	p.nombre = nuevoNombre
	fmt.Printf("Ahora %v tiene %v años.", p.nombre, p.edad)
}

func main() {

	p1 := persona{"Marcelo", 30}
	p1.getNombre()

}
