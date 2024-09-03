//Реализовать пересечение двух неупорядоченных множеств.

package main

import "fmt"

func intersection(num1, num2 []int) []int {
	//создаем мапу, где ключами будут элементы из 1 множества
	elem := make(map[int]bool)
	for _, i := range num1 {
		elem[i] = true
	}

	//создаем салйс для хранения пересечения
	var result []int
	for _, i := range num2 {
		//проверка на существование i в мапе elem
		if elem[i] {
			result = append(result, i)
			//что избежать дублирвания - удаляем найденное пересечение
			delete(elem, i)
		}
	}
	return result
}

func main() {
	//создаем 2 множества
	num1 := []int{2, 9, 4, 3, 8}
	num2 := []int{9, 5, 7, 2, 0}

	result := intersection(num1, num2)

	fmt.Println("Intersection:", result)

}
