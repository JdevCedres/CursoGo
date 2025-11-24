package main

import (
	"fmt"
)

func main() {
	// Operadores
	yearOld := 32

	fmt.Println("Operadores")
	fmt.Println(yearOld > 30)  // true
	fmt.Println(yearOld < 30)  // false
	fmt.Println(yearOld >= 40) // true
	fmt.Println(yearOld <= 32) // true
	fmt.Println(yearOld == 32) // true
	fmt.Println(yearOld != 32) // false

	fmt.Println()
	fmt.Println("Operadores Lógicos")
	fmt.Println("Operador OR")

	fmt.Println(yearOld < 32 || yearOld == 32) // (false || true) = true
	fmt.Println(yearOld < 32 || yearOld == 33) // (false || false) = false
	fmt.Println(yearOld < 40 || yearOld == 32) // (true || true) = true

	fmt.Println()
	fmt.Println("Operadores Lógicos")
	fmt.Println("Operador AND")

	fmt.Println(yearOld < 32 && yearOld == 32) // (false || true) = false
	fmt.Println(yearOld < 32 && yearOld == 33) // (false || false) = false
	fmt.Println(yearOld < 40 && yearOld == 32) // (true || true) = true

	fmt.Println()
	fmt.Println("Operadores Lógicos")
	fmt.Println("Operador NOT")

	fmt.Println(true)            // true
	fmt.Println(!true)           // false
	fmt.Println(yearOld < 40)    // true
	fmt.Println(!(yearOld < 40)) // false

	// Agrupar operadores con paréntesis
	fmt.Println()
	fmt.Println("Agrupar operadores con paréntesis")

	fmt.Println((yearOld < 32 && yearOld == 32) || (yearOld < 40 && yearOld == 32))
}
