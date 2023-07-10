package main

import "fmt"

type Human struct {
	name string
	age  int
}

func (h *Human) Speak() {
	fmt.Println("I am", h.name)
}

type Action struct {
	Human  // Встравание структуры
	action string
}

//...
func main() {
	action := Action{Human{"John", 30}, "running"}

	action.Speak() // вызов метода Speak из структуры Human, встроенной в структуру Action
	fmt.Println("I am", action.name, "and I am", action.age, "years old. I am", action.action)
}
