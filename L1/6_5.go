//ЗАВЕРШЕНИЕ ГЛОБАЛЬНОЙ ПЕРЕМЕННОЙ

package main

import (
	"fmt"
	"sync"
	"time"
)

var flag bool
var mu sync.Mutex //создание мьютекса

func worker() {
	for {
		mu.Lock() //блокировка мьютекса, значит доступ к глобальной
		// переменной может получить только данная горутина остальные ожидают
		if flag {
			mu.Unlock() //при true разблочим доступ и завершим программу
			fmt.Println("Goroutine stopped by flag")
			return
		}
		//при else разблокируем, так как далее не работаем с flag
		mu.Unlock()
		fmt.Println("Working...")
		time.Sleep(2 * time.Second)
	}
}

func main() {
	go worker()

	time.Sleep(5 * time.Second)

	//блокируем мьютекс перед изменением flag
	mu.Lock()
	flag = true
	//разблочим после изменения
	mu.Unlock()

	time.Sleep(2 * time.Second)
}
