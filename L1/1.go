package main

import "fmt"

//Дана структура Human (с произвольным набором полей и методов).
//Реализовать встраивание методов в структуре Action от
//родительской структуры Human (аналог наследования).

// создаем родительскую структуру Human
type Human struct {
	name string
	age  int
}

func (h Human) speak() {
	fmt.Printf("My name is %s and I am %d\n", h.name, h.age)
}

func (h Human) dance() {
	fmt.Println("I am dancing")
}

// встраиваем структуру Human в структуру Action,
// что позволяет напрямую вызывать методы Human из Action
type Action struct {
	Human
	ActionType string
}

// дополнительный метод Action,
// который может использовать поля и методы Human
func (a Action) PerformAction() {
	fmt.Printf("%s is performing an action: %s\n", a.name, a.ActionType)
}

func main() {
	person := Action{
		Human:      Human{name: "Kate", age: 20},
		ActionType: "Running",
	}

	person.speak()
	person.dance()
	person.PerformAction()
}
