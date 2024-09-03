//Дана последовательность температурных колебаний:
//-25.4, -27.0 13.0, 19.0, 15.5, 24.5, -21.0, 32.5.
//Объединить данные значения в группы с шагом в 10 градусов.
//Последовательность в подмножноствах не важна.

package main

import (
	"fmt"
	"math"
)

func main() {

	arr := []float64{-25.4, -27.0, 13.0, 19.0, 15.5, 24.5, -21.0, 32.5}
	group := make(map[int][]float64)

	for _, i := range arr {
		temp := int(math.Floor(i/10.0)) * 10
		group[temp] = append(group[temp], i)
	}

	for groups, temps := range group {
		fmt.Printf("Группа %d до %d: %v\n", groups, groups+10, temps)
	}
}
