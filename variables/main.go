package main

import (
	"fmt"
	"strconv"
	"unsafe"
)

func main() {
	var myIntVar int
	myIntVar = -12
	fmt.Println("My variable is: ", myIntVar)

	var myUintVar uint
	myUintVar = 12 // solo int positivo
	fmt.Println("My variable is: ", myUintVar)

	var myStringVar string
	myStringVar = "My string variable"
	fmt.Println("My variable is: ", myStringVar)

	var myBoolVar bool
	myBoolVar = true
	fmt.Println("My variable is: ", myBoolVar)

	// Otra manera de instanciar una variable :=

	myIntVar2 := 12
	fmt.Println("My variable myIntVar2 is:", myIntVar2)

	// Constantes

	const myFirstConst = "a12"
	fmt.Println("My const is: ", myFirstConst)

	const myIntConst int = 12
	fmt.Println("My const is: ", myIntConst)

	// Int

	fmt.Println()

	var my8BitsIntVar int8
	fmt.Printf("Int default value: %d\n", my8BitsIntVar)
	fmt.Printf("type: %T, bytes: %d, bits:%d\n", my8BitsIntVar, unsafe.Sizeof(my8BitsIntVar), unsafe.Sizeof(my8BitsIntVar)*8)
	var my16BitsIntVar int16
	fmt.Printf("type: %T, bytes: %d, bits:%d\n", my16BitsIntVar, unsafe.Sizeof(my16BitsIntVar), unsafe.Sizeof(my16BitsIntVar)*8)
	var my32BitsIntVar int32
	fmt.Printf("type: %T, bytes: %d, bits:%d\n", my32BitsIntVar, unsafe.Sizeof(my32BitsIntVar), unsafe.Sizeof(my32BitsIntVar)*8)
	var my64BitsIntVar int64
	fmt.Printf("type: %T, bytes: %d, bits:%d\n", my64BitsIntVar, unsafe.Sizeof(my64BitsIntVar), unsafe.Sizeof(my64BitsIntVar)*8)

	//uint también tiene los mismo 8, 16, 32, 64 bits solo positivos

	// Floats
	fmt.Println()

	var myFloat32var float32
	var myFloat64var float64
	fmt.Printf("Float default value: %f\n", myFloat32var)
	fmt.Printf("type: %T, bytes: %d, bits:%d\n", myFloat32var, unsafe.Sizeof(myFloat32var), unsafe.Sizeof(myFloat32var)*8)
	fmt.Printf("type: %T, bytes: %d, bits:%d\n", myFloat64var, unsafe.Sizeof(myFloat64var), unsafe.Sizeof(myFloat64var)*8)

	// String
	fmt.Println()

	var myStringVar3 string
	fmt.Printf("String default value: %s\n", myStringVar3)

	// Utilizar ` ` comillas simples hace saltos de lineas con el enter

	myStringvar4 := `My string variable in golang
	whith multiple 
	line`
	fmt.Printf("The variable value is: %s\n", myStringvar4)

	// Conversiones

	{
		fmt.Println()
		floatVar := 33.11
		fmt.Printf("type: %T, value: %f\n", floatVar, floatVar)
		floatStrVar := fmt.Sprintf("%.2f", floatVar) // Utilizamos Sprintf -> para pasar en este ejemplo de float a string
		fmt.Printf("type: %T, value: %s\n", floatStrVar, floatStrVar)

		intVar := 22
		fmt.Printf("type: %T, value: %d\n", intVar, intVar)
		intStrVar := fmt.Sprintf("%d", intVar)
		fmt.Printf("type: %T, value: %s\n", intStrVar, intStrVar)

		// Hacer conversión de un string a un int utilizamos package importado "strconv"

		intVal1, err := strconv.ParseInt("1234", 0, 64) // -> primero lo que queremos convertir en int , segundo desde donde empieza y cuanto de grade es el size en este caso 64 bits
		fmt.Println(err)
		fmt.Printf("type: %T, value: %d\n", intVal1, intVal1)
	}

}
