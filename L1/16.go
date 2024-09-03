//Реализовать быструю сортировку массива (quicksort)
//встроенными методами языка.

package main

import (
	"fmt"
	"sort"
)

func quicksort(arr []int) []int {
	if len(arr) < 2 {
		return arr
	}
	//опорный элемент
	ot := arr[0]

	var (
		less    []int
		greater []int
	)

	for _, num := range arr[1:] { //перебираем все элемнеты, кроме 1го
		if num <= ot {
			less = append(less, num)
		} else {
			greater = append(greater, num)
		}
	}

	return append(append(quicksort(less), ot), quicksort(greater)...)
}

func main() {
	arr := []int{10, 5, 2, 3, 7, 6, 4, 1, 9, 8}
	fmt.Println("My function:", quicksort(arr))

	//2 способ, использование sort.Ints
	arr = []int{10, 5, 2, 3, 7, 6, 4, 1, 9, 8}
	sort.Ints(arr)
	fmt.Println("sort.Ints:", arr)

}
