//Разработать программу, которая переворачивает подаваемую
//на ход строку (например: «главрыба — абырвалг»).
//Символы могут быть unicode.

package main

import "fmt"

func reverse(str string) string {
	result := []rune(str)
	lenght := len(str)
	for i := 0; i < lenght/2; i++ {
		result[i], result[lenght-1-i] = result[lenght-1-i], result[i]
	}
	return string(result)
}

func main() {
	var str string
	fmt.Scan(&str)

	fmt.Println(reverse(str))
}
