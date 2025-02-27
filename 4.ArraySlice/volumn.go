package main

import "fmt"

func main() {
	// 초기 용량을 5로 설정
	slice := make([]int, 0, 5)

	for i := 0; i < 10; i++ {
		slice = append(slice, i)
		fmt.Printf("Length: %d, Capacity: %d\n", len(slice), cap(slice))
	}
}
