package main

import "fmt"

func main() {
	day := 2

	switch day {
	case 1:
		fmt.Println("월요일")
	case 2:
		fmt.Println("화요일")
	case 3:
		fmt.Println("수요일")
	default:
		fmt.Println("주말")
	}
}