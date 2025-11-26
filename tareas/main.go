package main

import (
	"fmt"
)

func main() {
	fmt.Println("Tarea  1")

	license := false
	age := 19

	if license && age >= 15 {
		fmt.Println("Puede seguir avanzando")
	} else if !license || age <= 15 {
		fmt.Println("No puede seguir circulando siempre y cuando no tenga licencia o sea menor de edad")

	}

	fmt.Println("Tarea  2")
	array := [5]int{4, 2, 5, 6, 7}
	for i := range array {
		slice1 := array[i]
		fmt.Println(slice1 + 20)
	}

}
