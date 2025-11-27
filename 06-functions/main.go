package main

import (
	"fmt"

	"github.com/JdevCedres/CursoGo/06-functions/function"
)

func main() {

	fmt.Println(function.Add(3, 4))
	function.RepeatString(3, "Hello Go!")
	result, err := function.Calc(function.MUL, 2, 4)
	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Println("Value: ", result)
	}

	result, err = function.Calc(function.DIV, 10, 2)
	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Println("Value: ", result)
	}

	result, err = function.Calc(function.DIV, 2, 0)
	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Println("Value: ", result)
	}

	x, y := function.Split(20)
	fmt.Println("X:", x, "Y:", y)

	fmt.Println("MSum:", function.MSum(10, 20, 35, 47, 75, 68, 87, 88, 89, 100))

}
