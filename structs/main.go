package main

import (
	"fmt"
)

type persona struct {
	nombre string
	edad   int
}

func main() {

	p1 := persona{"Marcelo", 20}
	getNombre(p1)

}

func getNombre(p persona) {
	fmt.Printf("%v tiene %v años.", p.nombre, p.edad)
}
