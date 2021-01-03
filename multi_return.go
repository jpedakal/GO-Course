package main

import "fmt"

func main() {
	a, b := vals(5, 5)
	fmt.Println(a)
	fmt.Println(b)

}

func vals(a, b int) (int, int) {
	addNum := a + b
	mulNum := a * b
	return addNum, mulNum
}