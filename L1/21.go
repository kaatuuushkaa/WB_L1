//Реализовать паттерн «адаптер» на любом примере.

package main

import "fmt"

// класс собака
type Dog struct{}

// реакция собаки
func (dog *Dog) WoofWoof() {
	fmt.Println("Гав-гав. Хозяйка, дай поесть")
}

// класс кошка
type Cat struct{}

// реакция кошки, если ее не позвать - она молчит
func (cat *Cat) MeowMeow(isCall bool) {
	if isCall {
		fmt.Println("Где моя еда, раб? Мяу-мяу")
	}
}

// целевой  интерфейс
type AnimalAdapter interface {
	Reaction()
}

// адаптер собаки
type DogAdapter struct {
	*Dog
}

// реакция собаки:адаптер
func (adapter *DogAdapter) Reaction() {
	adapter.WoofWoof()
}

// конструкция адаптера для собаки
func NewDogAdapter(dog *Dog) AnimalAdapter {
	return &DogAdapter{dog}
}

// адаптер для кошки
type CatAdapter struct {
	*Cat
}

// реакция кошки:адаптер
func (adapter *CatAdapter) Reaction() {
	//адаптер автоматически зовет кошку
	adapter.MeowMeow(true)
}

// конструктор адаптер для кота
func NewCatAdapter(cat *Cat) AnimalAdapter {
	return &CatAdapter{cat}
}

func main() {
	fmt.Println("Установка адаптер с двуми чипами")
	myFamily := [2]AnimalAdapter{NewDogAdapter(&Dog{}), NewCatAdapter(&Cat{})}

	fmt.Println("Открываете дверь и заходите домой")
	for _, member := range myFamily {
		member.Reaction()
	}
}
