package main

import (
	"fmt"
)

func main() {
	DetermineType(1)
	DetermineType(1.1)
	DetermineType("1.1")
}

func DetermineType(value interface{}) {
	switch value.(type) {
	case int:
		fmt.Println("int")
	case float32:
		fmt.Println("float32")
	case float64:
		fmt.Println("float64")
	case string:
		fmt.Println("string")
	case bool:
		fmt.Println("bool")
	case chan interface{}:
		fmt.Println("chan")
	default:
		fmt.Println("unknown type")
	}
}
