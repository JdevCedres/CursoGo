package main

import (
	"fmt"
)

func main() {
	sum := 0
	for i := 1; i <= 10; i++ {
		sum++
	}
	fmt.Println(sum)

	sum = 1

	for sum < 1000 {
		sum++
	}
	fmt.Println(sum)

	for {
		if sum > 1000 {
			break
		}
		sum++
	}
	fmt.Println(sum)

	// recorrer un array
	arr := []int{1, 2, 3, 1, 2, 3}
	for i := range arr {
		fmt.Println("Indice:", i, "Valor:", arr[i])
	}

	// otro ejemplo recorrer array
	fmt.Println()
	for i, v := range arr { // si no nos hace falta el indice, podemos usar _
		fmt.Println("Indice:", i, "Valor:", v)
	}
	// recorrer un mapa
	fmt.Println()
	m2 := map[string]float64{
		"A": 12.3,
		"B": 23.1,
		"C": 34,
	}
	for k, v := range m2 {
		fmt.Println("Clave:", k, "Valor:", v)
	}

	// recorrer un array de maps
	fmt.Println()
	map3 := map[string][]int{
		"A": nil,
		"B": {2, 34, 1, 2, 4},
		"C": {4, 5, 3, 2, 1},
	}
	for key, value := range map3 {
		fmt.Println("key:", key)
		for _, v := range value {
			fmt.Println("   value:", v)
		}
	}
}
