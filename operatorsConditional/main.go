package main

import (
	"fmt"
)

func main() {
<<<<<<< HEAD
	yearsOld := 32

	fmt.Println("Operators")
	fmt.Println(yearsOld > 30)  // me va a devolver true
	fmt.Println(yearsOld < 32)  // me va a devolver false
	fmt.Println(yearsOld >= 32) // me va a devolver true
	fmt.Println(yearsOld <= 40) // me va a devolver false
	fmt.Println(yearsOld == 32) // me va a devolver true

	fmt.Println("Logical Operators")
	/*
		OR -> ||
		AND -> &&
		NOT -> !
	*/

	fmt.Println("OR")

	fmt.Println(yearsOld < 32 || yearsOld == 32) // (true and true) = true
	fmt.Println(yearsOld < 32 || yearsOld == 33) // (true and falsa) = false
	fmt.Println(yearsOld < 40 || yearsOld == 32) // (false and true) = true

	fmt.Println("AND")

	fmt.Println(yearsOld < 32 && yearsOld == 32) // (true and true) = true
	fmt.Println(yearsOld < 32 && yearsOld == 33) // (true and falsa) = false
	fmt.Println(yearsOld < 40 && yearsOld == 32) // (false and true) = false

	fmt.Println("NOT")

	fmt.Println(true)  // true
	fmt.Println(!true) // false

	fmt.Println(yearsOld < 40)    // True
	fmt.Println(!(yearsOld < 40)) // False

	fmt.Println()

	fmt.Println(yearsOld < 25 && yearsOld == 32 || yearsOld < 40)   // (false and true or true) = true
	fmt.Println(yearsOld < 25 && (yearsOld == 32 || yearsOld < 40)) // en este caso primero se hace lo que está dentro de los () y luego la siguiente. En este caso es falso.

	fmt.Println()
	fmt.Println("Condicionales IF")

	yearsOld = 20
	if yearsOld > 18 {
		fmt.Printf("%d is higher than 18\n", yearsOld)
	}

	boolVar := true
	if boolVar {
		fmt.Println("The variable is true")
	} else {
		fmt.Println("The variable is false")
	}

	number := 0
	switch number {
	case 1:
		fmt.Println("The number is one")
	case 2:
		fmt.Println("The number is two")
	case 3:
		fmt.Println("The number is three")
	default:
		fmt.Println("The number is unknown")
	}

	switch number {
	case 1:
		fmt.Println("The number is one")
	case 2:
		fmt.Println("The number is two")
	case 3:
		fmt.Println("The number is three")
	default:
		fmt.Println("The number is unknown")
	}

	switch {
	case number < 0:
		fmt.Println("The number is negative")
	case number > 0:
		fmt.Println("The number is positive")
	default:
		fmt.Println("The number is zero")
	}
=======
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
>>>>>>> c527ce1d42c2f6191849e483850b9e3f64d06417
}
