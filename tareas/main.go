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
		fmt.Println("No puede seguir circulando")

	}
}
