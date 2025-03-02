package main

import "fmt"

func main() {
	// 맴 정의
	myMap := make(map[string]int)

	// 데이터 추가
	myMap["apple"] = 1
	myMap["banana"] = 2

	// 데이터 조회
	fmt.Println("Apple:", myMap["apple"]); // result: Apple: 1

	// 삭제 후 조회
	value, exist := myMap["banana"]
	if exist {
		fmt.Println("Banana:", value); // result: Banana: 2
	} else {
		fmt.Println("Banana is not exist"); // result: Banana key dose not exist.
	}
	
	// 데이터 삭제
	delete(myMap, "apple"); // "apple" 키와 그에 해당하는 값 삭제
	fmt.Println("After deletion: ", myMap);

	// 전체 맵 출력
	for key, value := range myMap {
		fmt.Printf("%s: %d\n", key, value);
	}
}