package main

import (
	"fmt"
	"regexp"
)

func main() {

	match, _ := regexp.MatchString("p([a-z]+)ch", "peach")
	fmt.Println(match)

	// Above we used a string pattern directly,
	// but for other regexp tasks we'll need to Compile an optimized Regexp struct.

	r, _ := regexp.Compile("p([a-z]+)ch")
	fmt.Println(r.MatchString("peach"))
	fmt.Println(r.FindString("peach punch"))
	fmt.Println(r.FindStringIndex("peach punch"))
	fmt.Println(r.FindStringSubmatch("peach punch pinch"))
	fmt.Println(r.FindStringSubmatchIndex("peach punch"))
	fmt.Println(r.FindAllString("peach punch pinch", -1))
	fmt.Println("all:", r.FindAllStringSubmatchIndex(
		"peach punch pinch", -1))
	fmt.Println(r.FindAllString("peach punch pinch", 2))
	fmt.Println(r.ReplaceAllString("a peach", "<fruit>"))
}
