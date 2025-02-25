package main

import "fmt"


func main() { 
	func() {
		fmt.Println("Not Name Function")
	}()
}