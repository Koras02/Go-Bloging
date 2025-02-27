package main

import "fmt"

func main() {
	slice := []int{1,2,3} // 슬라이스 선언
	fmt.Println(slice); // [1,2,3]

	slice = append(slice, 4); // 슬라이스에 값 추가
	fmt.Println(slice) // [1 2 3 4]
}