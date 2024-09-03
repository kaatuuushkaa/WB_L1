//Реализовать собственную функцию sleep.

package main

import (
	"fmt"
	"time"
)

func Sleep(d time.Duration) {
	<-time.After(d) //блокировка канала на время d

}

func main() {
	fmt.Println("Do something...")
	Sleep(5 * time.Second)
	fmt.Println("A few moments later")
}
