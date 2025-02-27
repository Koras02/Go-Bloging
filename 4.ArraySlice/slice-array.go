package main

import "fmt"

func main() {
	var dynamicSlice []int // 슬라인스 선언 (nil 상태)

	// 슬라이스에 값 추가
	for i := 0; i < 10; i++ {
		dynamicSlice = append(dynamicSlice, i)
	}

	fmt.Println(dynamicSlice); // [0 1 2 3 4 5 6 7 8 9]

	// 슬라이스의 길이와 용량
	fmt.Println("Length:", len(dynamicSlice)) // Length: 10
	fmt.Println("Capacity:", len(dynamicSlice)) // Capacity: 10 (초기 용량)
}