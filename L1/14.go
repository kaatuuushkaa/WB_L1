//Разработать программу, которая в рантайме способна
//определить тип переменной: int, string, bool, channel
//из переменной типа interface{}.

package main

import "fmt"

func main() {
	var x interface{} = "1"

	switch expr := x.(type) {
	case int:
		fmt.Printf("int: %d", expr)
	case string:
		fmt.Printf("string: %s", expr)
	case bool:
		fmt.Printf("bool: %t", expr)
	case chan int:
		fmt.Printf("chan int")
	case chan string:
		fmt.Printf("chan string")
	case chan bool:
		fmt.Printf("chan bool")
	default:
		fmt.Println("I dont know")
	}
}
