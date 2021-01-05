package main

import (
	"fmt"
	s "strings"
)

var p = fmt.Println

func main() {
	p("contains:",s.Contains("test","es"))
	p("count ", s.Count("test","s"))
}
