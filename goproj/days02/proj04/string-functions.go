package main

import s "strings"
import "fmt"

var p = fmt.Println

func main() {

	p("Contains: ",s.Contains("test","s"))
	p("Count:    ",s.Count("test","t"))
	p("HasPrefix:",s.HasPrefix("test","te"))
	p("HasSuffix:",s.HasSuffix("test","te"))
	p("Index:    ",s.Index("test","e"))
	p("join:     ",s.Join([]string{"a","b","c"},"="))
	p("Repeat:   ",s.Repeat("a",5))
	p("Replace   ",s.Replace("foo","o","e",1))
	p("Split:    ",s.Split("f---o--o--","o"))
}
