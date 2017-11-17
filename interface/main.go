package main

import "fmt"

// Animal animal interface.
type Animal interface {
	Cry()
}

// Pet pet interface.
type Pet interface {
	Choker()
	LikeHuman()
}

// Dog struct
type Dog struct{}

// Cry dog cry method.
func (d *Dog) Cry() {
	fmt.Println("わんわん")
}

// Choker dog choker method.
func (d *Dog) Choker() {
	fmt.Println("リード付きの首輪")
}

// LikeHuman dog like human method.
func (d *Dog) LikeHuman() {
	fmt.Println("ご主人様大好き")
}

// Cat struct
type Cat struct{}

// Cry cat cry method
func (c *Cat) Cry() {
	fmt.Println("にゃーにゃー")
}

// Choker cat choker method.
func (c *Cat) Choker() {
	fmt.Println("鈴付きの首輪")
}

// LikeHuman cat like human method.
func (c *Cat) LikeHuman() {
	fmt.Println("ご主人様よりこたつ")
}

// Elephant elephant struct
type Elephant struct{}

// Cry elephant cry method.
func (e *Elephant) Cry() {
	fmt.Println("パオーン!")
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

// MakeSomeoneHasChoker choker func
func MakeSomeoneHasChoker(someone interface{}) {
	fmt.Println("どんな首輪してる？")
	p, ok := someone.(Pet)
	if !ok {
		fmt.Println("ペットじゃないから首輪してないよ")
		return
	}
	p.Choker()
}

func main() {
	dog := new(Dog)
	cat := new(Cat)
	stone := new(Stone)
	elephant := new(Elephant)
	MakeSomeoneCry(dog)
	MakeSomeoneCry(cat)
	MakeSomeoneCry(elephant)
	MakeSomeoneCry(stone)
	// MakeAnimalCry(dog)
	// MakeAnimalCry(cat)
	MakeSomeoneHasChoker(dog)
	MakeSomeoneHasChoker(cat)
	MakeSomeoneHasChoker(elephant)

	// h := Hex(10)
	// str := h.String()
	// fmt.Println(str)
}
