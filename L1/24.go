//Разработать программу нахождения расстояния между
//двумя точками, которые представлены в виде структуры Point с
//инкапсулированными параметрами x,y и конструктором.

package main

import (
	"fmt"
	"math"
)

type Point struct {
	x float64
	y float64
}

// конструктор для создания новой точки
func NewPoint(x, y float64) *Point {
	return &Point{x: x, y: y}
}

// вычисление расстояния между точками
func (p *Point) Distance(other *Point) float64 {
	dx := p.x - other.x
	dy := p.y - other.y
	return math.Sqrt(dx*dx + dy*dy)
}

func main() {
	p1 := NewPoint(1, 2)
	p2 := NewPoint(4, 6)

	distance := p1.Distance(p2)
	fmt.Println("Distance: ", distance)
}
