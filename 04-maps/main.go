package main

import (
	"fmt"
)

func main() {

	m1 := make(map[int]string)
	m1[1] = "A"
	m1[2] = "B"
	m1[3] = "C"

	fmt.Println(m1)
	fmt.Println(m1[1])

	// eliminar un valor delete
	delete(m1, 2)
	fmt.Println(m1)

	// modificar un valor
	m1[1] = "A2"
	fmt.Println(m1[1])

	// valores vacios y por defecto en este caso es un string vacio

	m1[7] = ""
	fmt.Println(m1[7])
	fmt.Println(m1[99])

	v, ok := m1[7]
	fmt.Println("The value:", v, "Present?", ok)

	v, ok = m1[99]
	fmt.Println("The value:", v, "Present?", ok)

	// Otra forma de declarar un mapa

	m2 := map[string]int{
		"A": 1,
		"B": 2,
		"C": 3,
	}
	fmt.Println(m2)

}
