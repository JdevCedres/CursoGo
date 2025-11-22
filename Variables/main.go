package main

import (
	"fmt"
	"unsafe"
)

func main() {

	// Variables

	var myIntVar int
	myIntVar = -12
	fmt.Println("My variable is: ", myIntVar)

	var myUintVar uint
	myUintVar = 12 // solo números positivos
	fmt.Println("My variable is: ", myUintVar)

	var myStringVar string
	myStringVar = "My string variable"
	fmt.Println("My variable is: ", myStringVar)

	var myBoolVar bool
	myBoolVar = true
	fmt.Println("My variable is: ", myBoolVar)

	fmt.Println("My string variable address is: ", &myStringVar) //ver la dirección de memoria de la variable.

	// Otra manera de instanciar una variable

	myIntVar2 := 12
	fmt.Println("My variable myIntVar2 is: ", myIntVar2)

	myStringVar2 := "My string variable whith := "
	fmt.Println("My string variable  2 is: ", myStringVar2)

	// Constantes

	const myFirstConst = "a12"
	fmt.Println("My const is: ", myFirstConst)

	const myIntConst int = 12
	fmt.Println("My const is: ", myIntConst)

	fmt.Println()

	var my8BitsIntVar int8
	fmt.Printf("Int default value: %d\n", my8BitsIntVar)

	fmt.Printf("type: %T, bytes: %d, bits: %d\n", my8BitsIntVar, unsafe.Sizeof(my8BitsIntVar), unsafe.Sizeof(my8BitsIntVar)*8)
}
