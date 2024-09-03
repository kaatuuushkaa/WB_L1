//Удалить i-ый элемент из слайса.

package main

import "fmt"

func main() {
	users := []string{"Bob", "Alice", "Kate", "Sam", "Tom"}
	// удаляем 3ий элемень
	var i = 2
	users = append(users[:i], users[i+1:]...)
	//users[:i] - сохраняем элементы до i
	// users[i+1:]... - добавляемэлементы после i
	fmt.Println(users)
}
