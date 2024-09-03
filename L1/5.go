//Разработать программу, которая будет последовательно
//отправлять значения в канал,
//а с другой стороны канала — читать.
//	По истечению N секунд программа должна завершаться.

package main

import (
	"fmt"
	"time"
)

// создаем функцию для генерации и записи значения в канал
func sendValue(ch chan int) {
	value := 1
	for {
		ch <- value
		value++
		time.Sleep(2 * time.Second)
	}
}

// функция, которая считывает и выводит значение из канала
func readValue(ch chan int) {
	for value := range ch {
		fmt.Printf("Value is %d\n", value)
	}
}

func main() {
	ch := make(chan int)

	go sendValue(ch)

	go readValue(ch)

	//врем работы программы
	n := 10
	time.Sleep(time.Duration(n) * time.Second)

	close(ch)
	fmt.Println("Program finished")
}
