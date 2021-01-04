package main

import (
	"crypto/sha1"
	"fmt"
)

func main() {

	s := "Welcome To DXC"
	h := sha1.New()
	h.Write([]byte(s))

	bs := h.Sum(nil)

	fmt.Printf("%x\n", bs)
}
