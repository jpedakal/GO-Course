package main

import (
	"fmt"
	"sort"
)

func main() {
	// strings sort
	strs := []string{"a", "b", "f", "c", "z", "q"}
	sort.Strings(strs)

	// ints sort
	digits := []int{1, 4, 6, 7, 78, 34, 56, 76, 12, 22}
	sort.Ints(digits)
	fmt.Println(digits)
}
