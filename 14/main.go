package main

import (
	"fmt"
	"reflect"
)

func main() {
	var variable interface{} = 512

	// Определение типа interface{}  switch
	switch variable.(type) {
	case int:
		fmt.Println("int")
	case string:
		fmt.Println("string")
	case bool:
		fmt.Println("bool")
	case chan int:
		fmt.Println("channel")
	default:
		fmt.Println("Unknown type")
	}

	// Использование рефлексии для определения типа переменной
	typeOfVariable := reflect.TypeOf(variable)
	fmt.Println("Type with reflection: ", typeOfVariable)
}
