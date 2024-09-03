//Разработать программу, которая перемножает,
//делит, складывает, вычитает две числовых переменных a,b,
//значение которых > 2^20.

package main

import (
	"fmt"
	"math/big"
)

func main() {
	a := new(big.Int)
	b := new(big.Int)

	a.SetString("10485760", 10) //(10^7)используем SetString вместо
	//big.NewInt, потому что его вместительность больше
	b.SetString("20971520", 10) //(2*10^7)

	fmt.Println("Sum: ", new(big.Int).Add(a, b)) //сложение
	fmt.Println("Sub: ", new(big.Int).Sub(a, b)) //вычитание
	fmt.Println("Mul: ", new(big.Int).Mul(a, b)) //умножение
	fmt.Println("Div: ", new(big.Int).Div(a, b)) //деление
	fmt.Println("Mod: ", new(big.Int).Mod(a, b)) //остаток от деления
}
