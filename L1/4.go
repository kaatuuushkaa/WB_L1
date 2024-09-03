// Реализовать постоянную запись данных в канал (главный поток).
//
//	Реализовать набор из N воркеров, которые читают произвольные
//
// данные из канала и выводят в stdout. Необходима возможность
// выбора количества воркеров при старте.
//
// Программа должна завершаться по нажатию Ctrl+C.
//
//	Выбрать и обосновать способ завершения работы всех воркеров.
package main

import (
	"fmt"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"
)

// функция воркера для чтения и вывода данных из канала
func worker(id int, jobs <-chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	for job := range jobs {
		fmt.Printf("Worker %d is working on %d job\n", id, job)
	}
	fmt.Printf("Worker %d stopping\n", id)
}

func main() {
	var numWorkers int
	fmt.Print("Enter number of workers: ")
	fmt.Scan(&numWorkers)

	//создаем канал
	jobs := make(chan int)

	var wg sync.WaitGroup

	for i := 1; i <= numWorkers; i++ {
		wg.Add(1)
		go worker(i, jobs, &wg)
	}

	//горутина для генерации задач
	go func() {
		jobID := 1
		for {
			//конструкция select позволяет ожидать несоклько операций с каналами,
			//в данном случае ожидает событие, которое произойдет через 1 сек
			select {
			case <-time.After(1 * time.Second):
				jobs <- jobID
				jobID++
			}
		}
	}()

	//создаем канал для обработки сигнала ctrl+c для завершения работы
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)

	//блокировка работы программы, до получения сигнала
	<-sigChan

	close(jobs)

	wg.Wait()

	fmt.Println("All workers have finished.")
}
