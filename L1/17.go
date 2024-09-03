//Реализовать бинарный поиск встроенными методами языка.

package main

import (
	"fmt"
	"sort"
)

func binarySearch(arr []int, target int) int {
	//задаем границы поиска
	left := 0
	right := len(arr) - 1

	for left <= right {
		//середина для дальнейшего разделения половин
		mid := left + (right-left)/2

		if arr[mid] == target {
			return mid
		}
		//ищем в левой половине
		if arr[mid] > target {
			right = mid - 1
		} else {
			//иначе в правой
			left = mid + 1
		}
	}
	return -1

}

func main() {
	arr := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	target := 4

	//отсортируем массив перед поиском
	sort.Ints(arr)

	index := binarySearch(arr, target)

	if index != -1 {
		fmt.Printf("Элемент %d найден на индексе %d.\n", target, index)
	} else {
		fmt.Printf("Элемент %d не найден в массиве.\n", target)
	}
}
