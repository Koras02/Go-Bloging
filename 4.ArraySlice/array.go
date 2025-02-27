package main

import "fmt"

func main() {
	var arr[6]int // 크기가 6일 배열 선언
	arr[0] = 1
	arr[1] = 2
	fmt.Println(arr) // [1,2,0,0,0,0]
}