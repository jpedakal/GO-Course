package main

import "fmt"

func main() {
	name := "jayakrishna"
	fmt.Println(&name)
	fmt.Println(*&name)
}
