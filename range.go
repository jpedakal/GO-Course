package main

import "fmt"

func main() {
	colorCodes := map[string]string{"white": "123", "green": "234", "red": "345"}
	for color, value := range colorCodes {
		fmt.Println(color, ":", value)
	}
}
