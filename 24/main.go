package main

import (
	"fmt"
	"math"
)

// Структура Point представляет точку с координатами x и y
type Point struct {
	x, y float64
}

// Конструктор для создания новой точки
func NewPoint(x, y float64) Point {
	return Point{x, y}
}

// Метод для вычисления расстояния между двумя точками
func (p1 Point) Distance(p2 Point) float64 {
	dx := p2.x - p1.x
	dy := p2.y - p1.y
	return math.Sqrt(dx*dx + dy*dy)
}

func main() {
	p1 := NewPoint(1, 2)
	p2 := NewPoint(4, 6)
	fmt.Println(p1, p2)

	// Вычисляем расстояние между ними
	distance := p1.Distance(p2)

	fmt.Println(distance)
}
