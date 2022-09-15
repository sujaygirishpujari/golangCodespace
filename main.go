package main

import "fmt"

func add(x, y int) int {
	return x + y
}

func main() {

	var confName = "abc"
	fmt.Printf("confName is %v"+" And addition of 5 + 7 = %v", confName, add(5, 7))
}
