package main

import "fmt"

func divide(a int, b int) (int, int) {
	return a / b, a % b
}

func main() {
	quotient, remainder := divide(10, 5)
	fmt.Println("result:", quotient, "나머지:", remainder)
}