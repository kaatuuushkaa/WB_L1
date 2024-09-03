//Реализовать структуру-счетчик, которая будет
//инкрементироваться в конкурентной среде.
//По завершению программа должна выводить
//итоговое значение счетчика.

package main

import (
	"fmt"
	"sync"
)

// структура-счетчик
type Counter struct {
	num int
	mu  sync.Mutex
}

func (c *Counter) Increment() {
	c.mu.Lock()
	c.num++
	c.mu.Unlock()
}

func main() {
	counter := Counter{}

	var wg sync.WaitGroup
	arr := []int{5, 2, 9, 13, 0, -12}
	for _, i := range arr {
		wg.Add(1)
		go func(a int) {
			defer wg.Done()
			counter.Increment() //внутри структуры счетчик увеличивается
		}(i)
	}
	wg.Wait()
	defer fmt.Printf("Итоговое значение счетчика %d", counter.num)

}
