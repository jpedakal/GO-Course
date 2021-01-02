package main

import "fmt"

func main() {
	res := addNum(100, 200)
	fmt.Println(res)
}

func addNum(a, b int) int {
	return a + b
}
