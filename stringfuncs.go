package main

import (
	"fmt"
	s "strings"
)

var p = fmt.Println // alias are possible, here for fmt.Println

func main() {
	p("Contains: ", s.Contains("test", "es"))
	p("Count: ", s.Count("test", "t"))
	p("HasPrefix: ", s.HasPrefix("test", "te"))
	p("HasSuffix: ", s.HasSuffix("test", "st"))
	p("Index: ", s.Index("test", "e"))
	b := []string{"goblin", "orc"}
	c := []string{"-", "are not friends"}
	d := s.Join(b, " & ")
	e := s.Join(c, " ")
	p("Join3: ", s.Join([]string{d, e}, " "))
	p("Repeat: ", s.Repeat("a", 5))
	p("Replace: ", s.Replace("foo", "o", "55", -1)) // 4th argument defines how many replacements to do
	p("Replace: ", s.Replace("foo", "o", "1", 1))
	p("Split: ", s.Split("This sentence should be stored in array as separate words", " "))
	p("ToLower: ", s.ToLower("UPPERCASE INITIALLY"))
	p("ToUpper: ", s.ToUpper("lowercase intially"))
}
