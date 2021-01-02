package main

import "fmt"

func main() {
	colors := []string{"white", "grey", "yellow", "green"}
	
	for i := 0; i < len(colors); i++ {
		fmt.Println(colors[i])
	}
}
