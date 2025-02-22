package main

// var 변수이름 데이터 타입
// var age int;

// 문자열(string)
var greeting string = "Hello,World!"

// 불리언(boolean)
var isActive bool = true

// 슬라이스(Slice)
var number []int = []int{1,2,3}

// 맵 (Map)
var ages map[string]int = map[string]int{"Konan": 25., "Bobb": 30}

// 구조체 (Struct)
type Person struct {
	Name string
	Age int 
	Job string 
}

// 인터페이스(interface)
type Animal interface {
	Speak() string
}

// 타입 추론 (Type Interface)
var age = 27 // int로 추론
name := "Jim" // string으로 추론론

// 명시적 타입 지정(Explicit Type Specification)
var height float64 = 1.82
var isActive bool = true