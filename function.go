package main

import "fmt"

func main() {
	res := addNum(100, 200)
	fmt.Println(res)
}

// type of arguements are int and return value type is int
func addNum(a, b int) int {
	return a + b
}
