//Дана переменная int64. Разработать программу
//которая устанавливает i-й бит в 1 или 0.

package main

import (
	"fmt"
)

func setBit(n int64, i uint, value bool) int64 {
	if value {
		n |= (1 << i)
		// "|" - битовое ИЛИ
		//"(1 << i)" - в i-тый бит устанавливаем 1
	} else {
		n &= ^(1 << i)
		//"^" - битовое НЕ(инверсия)
		//"&" - битовое И
		//
	}
	return n
}

func main() {
	var n int64 = 0
	fmt.Printf("Original: %064b\n", n)

	i := uint(3) //устанавливаем переменную для работы
	// с i-тым битом, в нашем случае 3им

	n = setBit(n, i, true)
	fmt.Printf("After setting bit %d to 1: %064b\n", i, n)

	n = setBit(n, i, false)
	fmt.Printf("After setting bit %d to 1: %064b\n", i, n)

}
