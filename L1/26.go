//Разработать программу, которая проверяет, что все символы в
//строке уникальные (true — если уникальные, false etc).
//Функция проверки должна быть регистронезависимой.
//
//Например:
//abcd — true
//abCdefAaf — false
//	aabcd — false

package main

import "fmt"

func IsUniq(str string) bool {
	tmp := make(map[rune]bool)

	for _, i := range str {
		if tmp[i] {
			return false
		}
		tmp[i] = true
	}
	return true
}

func main() {
	var str string
	fmt.Scan(&str)
	if IsUniq(str) {
		fmt.Println("True")
	} else {
		fmt.Println("False")
	}
}
