//ИСПОЛЬЗОВАНИЕ CONTEXT.CONTEXT

package main

import (
	"context"
	"fmt"
	"time"
)

func worker(ctx context.Context) {
	for {
		select {
		case <-ctx.Done(): //горутина, следящая за сигналом отмены
			fmt.Println("Goroutine stopped by context cancel")
			return
		default:
			fmt.Println("Working...")
			time.Sleep(2 * time.Second)
		}
	}
}

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	//context.Background - корневой контекст для всей программы
	//context.WithCancel(parent Context) - контекст, который можно отменить вручную

	go worker(ctx)

	time.Sleep(5 * time.Second)
	cancel() //отменяем контекст, что приведет к завершения горутины
	time.Sleep(2 * time.Second)

}
