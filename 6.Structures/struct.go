package main

import "fmt"

// Animal 구조체 정의
type Animal struct {
	Name string 
	Age int 
}

// interface 정의
type Speaker interface {
	Speak() string
}

// 사용자 정의 데이터 타입
type Dog struct {
	Animal
}

type Cat struct {
	Animal
}

// Speak 메서드 구현
func (a Dog) Speak() string {
	return "Woof!"
}

func (a Cat) Speak() string {
	return "Meow!"
}

func makeSound(s Speaker) {
	fmt.Println(s.Speak())
}

func main() {
	dog := Dog{Animal{"Buddy", 3}}
	cat := Cat{Animal{"Whiskers", 2}}

	makeSound(dog) // result: Woof!
	makeSound(cat) // result: Meow!
}