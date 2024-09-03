//Реализовать все возможные способы
//остановки выполнения горутины.

//ЗАВЕРШЕНИЕ ПО КАНАЛУ(СИГНАЛ ЗАВЕРШЕНИЯ)

package main

import (
	"fmt"
	"time"
)

// функция воркер, которая завершает работу при получении сигнала в канал
func worker(done chan bool) {
	for {
		select {
		case <-done:
			fmt.Println("Goroutine stopped by signal")
			return
		default:
			fmt.Println("Working...")
			time.Sleep(1 * time.Second)
		}
	}
}

func main() {
	done := make(chan bool)

	go worker(done)
	time.Sleep(2 * time.Second)
	done <- true                //отправляем сигнал для остановки программы
	time.Sleep(1 * time.Second) //даем время для завершения горутины
}
