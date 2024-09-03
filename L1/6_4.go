//ИСПОЛЬЗОВАНИЕ ЗАКРЫТИЯ КАНАЛА

package main

import (
	"fmt"
	"time"
)

func worker(ch chan struct{}) {
	for {
		select {
		case <-ch:
			fmt.Println("Goroutine stopped by closing channel")
			return
		default:
			fmt.Println("Working...")
			time.Sleep(2 * time.Second)
		}
	}
}

func main() {
	ch := make(chan struct{})

	go worker(ch)

	time.Sleep(5 * time.Second)
	close(ch)
	time.Sleep(2 * time.Second)
}
