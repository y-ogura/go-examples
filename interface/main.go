package main

import "fmt"

// Animal animal interface.
type Animal interface {
	Cry()
}

// Dog struct
type Dog struct{}

// Cry dog cry method.
func (d *Dog) Cry() {
	fmt.Println("わんわん")
}

// Cat struct
type Cat struct{}

// Cry cat cry method
func (c *Cat) Cry() {
	fmt.Println("にゃーにゃー")
}

// Stone struct
type Stone struct{}

// MakeAnimalCry cry func
func MakeAnimalCry(a Animal) {
	fmt.Println("鳴け!")
	a.Cry()
}

// MakeSomeoneCry cry func
func MakeSomeoneCry(someone interface{}) {
	fmt.Println("鳴け!")
	a, ok := someone.(Animal)
	if !ok {
		fmt.Println("動物ではないので鳴けません。")
		return
	}
	a.Cry()
}

func main() {
	dog := new(Dog)
	cat := new(Cat)
	stone := new(Stone)
	MakeSomeoneCry(dog)
	MakeSomeoneCry(cat)
	MakeSomeoneCry(stone)
	// MakeAnimalCry(dog)
	// MakeAnimalCry(cat)

	// h := Hex(10)
	// str := h.String()
	// fmt.Println(str)
}
