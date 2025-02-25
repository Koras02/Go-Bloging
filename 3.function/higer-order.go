package main

import "fmt"

func add (a int, b int) int {
	return a + b
}

func apply(f func(int, int) int, x int, y int) int {
	return f(x, y)
}

func main() {
	result := apply(add, 2, 3) // 5
	fmt.Println(result) // 5
}