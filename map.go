package main

import "fmt"

func main() {
	colorCodes := map[string]string{"white": "123", "green": "234", "red": "345"}
	colorCodes["yellow"] = "567"
	fmt.Println(colorCodes)
}
