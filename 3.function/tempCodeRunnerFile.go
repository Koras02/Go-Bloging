package main

import "fmt"

func makeCounter() func() int {
	count := 0

	return func() int {
		count++ 
		return count
	}
}

func main() {
	counter := makeCounter()
	fmt.Println(counter()) // 1
	fmt.Println(counter()) // 2

}