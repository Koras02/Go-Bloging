package main

import "fmt"

func main() {
	// 변수 선언 초기화
	var age int = 25
	name := "Jick"
	isActive := true 

	// 타입추론
	var height = 1.85 // float64

	// 명시적 타입 지정
	var score float32 = 95.5

	// 출력
	fmt.Println("Name:", name)
	fmt.Println("Age:", age)
	fmt.Println("Active:", isActive)
	fmt.Println("Height:", height)
	fmt.Println("Score:", score)

}