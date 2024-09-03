//использование sync.Map

package main

import (
	"fmt"
	"sync"
)

func main() {
	//больше подойдет при частом
	//чтении и редком записывании
	var myMap sync.Map

	var wg sync.WaitGroup

	writeToMap := func(key int, value string) {
		defer wg.Done()
		myMap.Store(key, value)
	}

	for i := 0; i < 5; i++ {
		wg.Add(1)
		go writeToMap(i, fmt.Sprintf("Value %d", i))
	}

	wg.Wait()

	myMap.Range(func(key, value interface{}) bool {
		fmt.Printf("key: %d, value: %s\n", key, value)
		return true
	})

}
