package main

import "fmt"

type Animal interface {
	Cry()
}

type Dog struct{}

func (d *Dog) Cry() {
	fmt.Println("わんわん")
}

type Cat struct{}

func (c *Cat) Cry() {
	fmt.Println("にゃーにゃー")
}

func MakeAnimalCry(a Animal) {
	fmt.Println("鳴け!")
	a.Cry()
}

func main() {
	dog := new(Dog)
	cat := new(Cat)
	MakeAnimalCry(dog)
	MakeAnimalCry(cat)

	var h Hex
	h = 10
	h.String()
}
