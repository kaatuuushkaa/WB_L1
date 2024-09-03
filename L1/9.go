//Разработать конвейер чисел. Даны два канала: в первый пишутся числа (x) из массива,
//во второй — результат операции x*2,
//после чего данные из второго канала должны выводиться в stdout.

package main

import "fmt"

func main() {
	//создаем 2 канала
	numbers := make(chan int)
	result := make(chan int)

	//в 1 канал вписываем числа
	go func() {
		arr := []int{1, 2, 3, 4, 5}
		for _, x := range arr {
			numbers <- x
		}
		close(numbers)
	}()

	//вычисляем х*2 и записываем во 2 канал
	go func() {
		for x := range numbers {
			result <- x * 2
		}
		close(result)
	}()

	//выводим результат из 2 канала
	for num := range result {
		fmt.Println(num)
	}
}
