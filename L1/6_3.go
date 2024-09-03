//ЗАВЕРШЕНИЕ ПО ТАЙМИНГУ CONTEXT.CONTEXT

package main

import (
	"context"
	"fmt"
	"time"
)

func worker(ctx context.Context) {
	for {
		select {
		case <-ctx.Done():
			fmt.Println("Goroutine stopped by context timeout")
			return
		default:
			fmt.Println("Working...")
			time.Sleep(1 * time.Second)
		}
	}
}

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	//context.WithTimeout(parent Context, timeout Duration)-контекст
	//автоматически отменяется по истечению времени

	defer cancel()

	go worker(ctx)

	time.Sleep(3 * time.Second)
}
