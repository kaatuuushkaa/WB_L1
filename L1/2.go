package main

//Написать программу, которая конкурентно рассчитает значение квадратов
//чисел взятых из массива (2,4,6,8,10)
//и выведет их квадраты в stdout.

import (
	"fmt"
	"sync"
)

// чтобы результат выводился последовательно, я выбрала работу с каналами и структурой,
// в которую будет сохраняться результат
type Result struct {
	index int
	value int
}

func square(num, index int, wg *sync.WaitGroup, results chan<- Result) {
	defer wg.Done()
	//после завершения вычисления уменьшаем счетчик на 1
	// и уведомляем WaitGroup о том, что 1 горутина завершена
	results <- Result{index: index, value: num * num}
}

func main() {
	numbers := []int{2, 4, 6, 8, 10}
	var wg sync.WaitGroup
	//переменная для синхронизации
	// выполнения горутин, позволяет ожидать завершения
	//выполнения для всех горутин
	results := make(chan Result, len(numbers))

	for i, num := range numbers {
		wg.Add(1)
		// перед запуском горутины увеличиваем счетчик
		//и теперь WaitGroup будет знать о каждой запущенной горутине
		go square(num, i, &wg, results)
	}

	wg.Wait()
	//блокирует выполнение до тех пор, пока счетчик не будет равен 0
	//сообщаем WaitGroup, что в программе появилось новое
	// параллельное задание, за завершением которого надо следить, до
	//завершения всех запущенных горутин
	close(results)

	squares := make([]int, len(numbers))
	for result := range results {
		squares[result.index] = result.value
	}

	for i, square := range squares {
		fmt.Printf("Square of %d is %d\n", numbers[i], square)
	}
}
