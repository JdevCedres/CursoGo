package main

import (
	"fmt"
	"unsafe"
)

func main() {

	// Vectores / Arrays tienen un tamaño fijo

	var myIntVar int
	fmt.Println(myIntVar)
	fmt.Printf("type: %T, bytes: %d, bits: %d\n", myIntVar, unsafe.Sizeof(myIntVar), unsafe.Sizeof(myIntVar)*8)

	var myArrayVar [6]int
	fmt.Println("Array:", myArrayVar)

	myArrayVar[1] = 2
	myArrayVar[2] = 5
	myArrayVar[3] = 9
	fmt.Println("Array modified:", myArrayVar)
	fmt.Println("Size: ", len(myArrayVar))
	fmt.Printf("type: %T, bytes: %d, bits: %d\n", myArrayVar, unsafe.Sizeof(myArrayVar), unsafe.Sizeof(myArrayVar)*8)

	myArrayVar1 := [3]string{"value1", "value2", "value3"}
	fmt.Println(myArrayVar1)

	// Slice - no tienen tamaño fijo, tinen un tamaño flexible y dinamico

	var mySliceVar []int
	fmt.Printf("size: %d, value %v\n", len(mySliceVar), mySliceVar)

	mySliceVar = append(mySliceVar, 10, 20, 30, 40, 50)
	fmt.Printf("size: %d, value %v\n", len(mySliceVar), mySliceVar)

	mySlice := myArrayVar[1:4] // desde la posición 1 hasta la 4 (sin incluir la 4)
	fmt.Println(mySlice)
	fmt.Println("Sice: ", len(mySlice))

	fmt.Println(&myArrayVar[2])
	fmt.Println(&mySlice[1])

	fmt.Println(myArrayVar)

	mySlice[0] = 80
	mySlice[1] = 100

	fmt.Println("Array modified after changing slice:", myArrayVar)

	fmt.Println(myArrayVar[:4])
	fmt.Println(myArrayVar[2:])

	slice := make([]int, 3)
	fmt.Println(slice)
	fmt.Printf("size: %d, value: %v\n", len(slice), slice)

}
