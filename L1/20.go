//Разработать программу, которая переворачивает слова в строке.
//Пример: «snow dog sun — sun dog snow».

package main

import (
	"fmt"
	"strings"
)

func main() {
	str := "snow dog sun"
	//делим строку на слова
	words := strings.Fields(str)
	var result []string
	for i := len(words) - 1; i >= 0; i-- {
		result = append(result, words[i])
	}
	fmt.Println(result)
}
