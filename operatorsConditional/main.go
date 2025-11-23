package main

import (
	"fmt"
)

func main() {
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

}
