//Имеется последовательность строк -
//(cat, cat, dog, cat, tree)
//создать для нее собственное множество.

package main

import "fmt"

func main() {
	arr := []string{"cat", "cat", "dog", "cat", "tree"}

	result := make(map[string]bool)
	for _, i := range arr {
		result[i] = true
	}

	fmt.Println("Уникальные строки(множество):")
	for i := range result {
		fmt.Println(i)
	}

}
