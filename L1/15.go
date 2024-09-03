//К каким негативным последствиям может привести данный фрагмент кода,
//и как это исправить? Приведите корректный пример реализации.
//
//
//var justString string
//func someFunc() {
//  v := createHugeString(1 << 10)
//  justString = v[:100]
//}
//
//func main() {
//  someFunc()
//}

//проблема: код держит в памяти целую
//большую строку, даже если нужно только ее небольшая часть,
//и может неправильно обрезать многобайтовые символы

//исправление: мы берем только нужные символы и создаем
//из них новую строку, что сохраняет память

package main

var justString string

func someFunc() {
	v := createHugeString(1 << 10) //создается большая строка size=1024
	runes := []rune(v)             //преобразуем строку в срез рун, чтобы коректно высчитать len
	if len(runes) > 100 {
		justString = string(runes[:100]) //создаем новую строку из первых 100 символов
	} else {
		justString = string(runes)
	}

}

func createHugeString(size int) string {
	return string(make([]rune, size))
}

func main() {
	someFunc()
}
