//Реализовать конкурентную запись данных в map.
//использование мьютекса

package main

import (
	"fmt"
	"sync"
)

func main() {
	myMap := make(map[int]string)

	//мьютекс для синхронизации доступа к мапе
	var mu sync.Mutex
	//waitGroup для ожидания завершения всех горутин
	//гарантирует, что основаня программа не завершится, пока не закончат работать все горутины
	var wg sync.WaitGroup

	writeToMap := func(key int, value string) {
		defer wg.Done()
		mu.Lock()
		myMap[key] = value
		mu.Unlock()
	}

	for i := 0; i < 5; i++ {
		wg.Add(1)
		go writeToMap(i, fmt.Sprintf("Value %d", i))
	}

	wg.Wait()

	for key, value := range myMap {
		fmt.Printf("key: %d, value: %s\n", key, value)
	}

}
