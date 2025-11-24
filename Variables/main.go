package main

import (
	"fmt"
	"strconv"
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
	fmt.Println("INT TYPES")

	var my8BitsIntVar int8
	fmt.Printf("Int default value: %d\n", my8BitsIntVar)
	fmt.Printf("type: %T, bytes: %d, bits: %d\n", my8BitsIntVar, unsafe.Sizeof(my8BitsIntVar), unsafe.Sizeof(my8BitsIntVar)*8)

	var my16BitsIntVar int16
	fmt.Printf("Int default value: %d\n", my16BitsIntVar)
	fmt.Printf("type: %T, bytes: %d, bits: %d\n", my8BitsIntVar, unsafe.Sizeof(my8BitsIntVar), unsafe.Sizeof(my16BitsIntVar)*8)

	var my32BitsIntVar int32
	fmt.Printf("Int default value: %d\n", my16BitsIntVar)
	fmt.Printf("type: %T, bytes: %d, bits: %d\n", my8BitsIntVar, unsafe.Sizeof(my8BitsIntVar), unsafe.Sizeof(my32BitsIntVar)*8)

	var my64BitsIntVar int64
	fmt.Printf("Int default value: %d\n", my16BitsIntVar)
	fmt.Printf("type: %T, bytes: %d, bits: %d\n", my8BitsIntVar, unsafe.Sizeof(my8BitsIntVar), unsafe.Sizeof(my64BitsIntVar)*8)

	fmt.Println()
	fmt.Println("FLOAT TYPES")

	var my32BitsFloatVar float32
	fmt.Printf("Float default value: %f\n", my32BitsFloatVar)
	fmt.Printf("type: %T, bytes: %d, bits: %d\n", my32BitsFloatVar, unsafe.Sizeof(my32BitsFloatVar), unsafe.Sizeof(my32BitsFloatVar)*8)

	var my64BitsFloatVar float64
	fmt.Printf("Float default value: %f\n", my64BitsFloatVar)
	fmt.Printf("type: %T, bytes: %d, bits: %d\n", my64BitsFloatVar, unsafe.Sizeof(my64BitsFloatVar), unsafe.Sizeof(my64BitsFloatVar)*8)

	fmt.Println()
	fmt.Println("UINT TYPES")

	var my8BitsUIntVar uint8
	fmt.Printf("Uint default value: %d\n", my8BitsUIntVar)
	fmt.Printf("type: %T, bytes: %d, bits: %d\n", my8BitsUIntVar, unsafe.Sizeof(my8BitsUIntVar), unsafe.Sizeof(my8BitsUIntVar)*8)

	var my16BitsUIntVar uint16
	fmt.Printf("Uint default value: %d\n", my16BitsUIntVar)
	fmt.Printf("type: %T, bytes: %d, bits: %d\n", my16BitsUIntVar, unsafe.Sizeof(my16BitsUIntVar), unsafe.Sizeof(my16BitsUIntVar)*8)

	var my32BitsUIntVar uint32
	fmt.Printf("Uint default value: %d\n", my32BitsUIntVar)
	fmt.Printf("type: %T, bytes: %d, bits: %d\n", my32BitsUIntVar, unsafe.Sizeof(my32BitsUIntVar), unsafe.Sizeof(my32BitsUIntVar)*8)

	var my64BitsUIntVar uint64
	fmt.Printf("Uint default value: %d\n", my64BitsUIntVar)
	fmt.Printf("type: %T, bytes: %d, bits: %d\n", my64BitsUIntVar, unsafe.Sizeof(my64BitsUIntVar), unsafe.Sizeof(my64BitsUIntVar)*8)

	fmt.Println()
	fmt.Println("STRING TYPE")

	var myStringVar3 string
	fmt.Printf("String default value: %s\n", myStringVar3)

	myStringVar4 := `My string in golang can span
multiple lines using backticks.`
	fmt.Printf("My string variable 4 is: %s\n", myStringVar4)

	// Conversiones de las variables

	{
		fmt.Println()
		fmt.Println("VARIABLE CONVERSIONS")
		floatVar := 33.11
		fmt.Printf("type: %T, value: %f\n", floatVar, floatVar)
		// Convertir float a string
		floatStrVar := fmt.Sprintf("%.2f", floatVar)
		fmt.Printf("type: %T, value: %s\n", floatStrVar, floatStrVar)

		// Convertir int a string
		intVar := 22
		fmt.Printf("type: %T, value: %d\n", intVar, intVar)
		intStrVar := fmt.Sprintf("%d", intVar)
		fmt.Printf("type: %T, value: %s\n", intStrVar, intStrVar)

		// Convertir string a int
		intVal1, err := strconv.ParseInt("121212", 0, 64)
		fmt.Println(err)
		fmt.Printf("type: %T, value: %d\n", intVal1, intVal1)

		floatVar1, _ := strconv.ParseFloat("-11.22", 64)
		fmt.Printf("type: %T, value: %.2f\n", floatVar1, floatVar1)

		// Tipo byte

		fmt.Println()
		fmt.Println("BYTE TYPE")

		var A byte = 'A'
		var a byte = 'a'
		var R byte = 82
		var s byte = 115
		vector := []byte{65, 97, 82, 115}

		fmt.Println("Value in ASCII Code:")
		fmt.Println(A)
		fmt.Println(a)
		fmt.Println(R)
		fmt.Println(s)
		fmt.Println(vector)

		fmt.Println("Value in String:")
		fmt.Println(string(A))
		fmt.Println(string(a))
		fmt.Println(string(R))
		fmt.Println(string(s))
		fmt.Println(string(vector))
	}
}
