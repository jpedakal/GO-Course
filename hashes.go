package main

import (
	"crypto/sha1"
	"fmt"
)

func main() {

	s := "Welcome To DXC"
	h := sha1.New()
	fmt.Printf("%x\n", h)
	h.Write([]byte(s))
	fmt.Printf("%x\n", h)
	bs := h.Sum(nil)
	fmt.Printf("%x\n", h)
	fmt.Println(s)

	fmt.Printf("%x\n", bs)

}
