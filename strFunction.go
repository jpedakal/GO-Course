package main

import (
	"fmt"
	s "strings"
)

var p = fmt.Println

func main() {
	str := "test"
	p("contains:",s.Contains("test","es"))
	p("count:", s.Count(str,"s"))
	p("HasPrefix:", s.HasPrefix(str,"te"))
	p("HasSuffix:", s.HasSuffix(str,"st"))
	p("Index:", s.Index(str,"e"))
	p("Join:", s.Join([]string{"a","b"},"-"))
	p("Repeat", s.Repeat("a",5))
}
