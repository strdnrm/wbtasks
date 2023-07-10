package main

import "fmt"

// Интерфейс для работы с фигурами
type Shape interface {
	area() float64
}

// Квадрат
type Square struct {
	sideLength float64
}

func (s *Square) area() float64 {
	return s.sideLength * s.sideLength
}

// Адаптер для работы с прямоугольником как с фигурой
type RectangleAdapter struct {
	rectangle *Rectangle
}

func (ra *RectangleAdapter) area() float64 {
	return ra.rectangle.length * ra.rectangle.width
}

// Прямоугольник
type Rectangle struct {
	length float64
	width  float64
}

// Функция, которая принимает фигуру и выводит ее площадь
func printArea(shape Shape) {
	fmt.Println("Area:", shape.area())
}

func main() {
	square := &Square{sideLength: 5}
	rectangle := &Rectangle{length: 4, width: 6}

	printArea(square)

	// Создаем адаптер для работы с прямоугольником как с фигурой
	rectangleAdapter := &RectangleAdapter{rectangle: rectangle}
	printArea(rectangleAdapter)
}
