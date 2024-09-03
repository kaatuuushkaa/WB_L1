//Дана последовательность чисел: 2,4,6,8,10.
//Найти сумму их квадратов(22+32+42….)
//с использованием конкурентных вычислений.

package main

import (
	"fmt"
	"sync"
)

func squareAndSum(num int, ch chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	ch <- num * num
}

func main() {
	numbers := []int{2, 4, 6, 8, 10}
	var wg sync.WaitGroup

	ch := make(chan int, len(numbers))
	sum := 0

	for _, num := range numbers {
		wg.Add(1)
		go squareAndSum(num, ch, &wg)
	}

	wg.Wait()
	close(ch)

	for val := range ch {
		sum += val
	}

	fmt.Printf("Sum of squares: %d\n", sum)
}
