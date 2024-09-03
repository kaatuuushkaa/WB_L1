//Поменять местами два числа без создания временной переменной.

package main

import "fmt"

func main() {
	a := 1
	b := 2
	fmt.Println("BEFORE")
	fmt.Println("a:", a)
	fmt.Println("b:", b)

	//с помощью сложения и вычитания
	a = a + b
	b = a - b
	a = a - b

	//с помощью умножения
	//a=a*b
	//b=a/b
	//a=a/b

	fmt.Println("AFTER")
	fmt.Println("a:", a)
	fmt.Println("b:", b)
}
