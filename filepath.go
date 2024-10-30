package main

import (
	"fmt"
	"path/filepath"
	"strings"
)

func main() {
	p := filepath.Join("dir1", "dir2", "filename")
	fmt.Println("p: ", p)

	fmt.Println(filepath.Join("dir1//", "filename"))
	fmt.Println(filepath.Join("dir1/../dir1", "filename"))

	fmt.Println("Dir(p): ", filepath.Dir(p))
	fmt.Println("Base(p): ", filepath.Base(p))
	fmt.Println(filepath.Split(p)) // filepath.Split splits path when it's just a string

	fmt.Println(filepath.IsAbs("dir/file"))
	fmt.Println(filepath.IsAbs("/dir/file")) // IsAbs checks if the path is absolute, here: true

	filename := "config.json"
	ext := filepath.Ext(filename)
	fmt.Println(ext) // file extension extraction with . included
	fmt.Println(strings.TrimSuffix(filename, ext))

	// Rel returns a relative path from base(arg1) to target(arg2)
	rel, err := filepath.Rel("a/b", "a/b/t/file")
	if err != nil {
		panic(err)
	}
	fmt.Println(rel)

	rel, err = filepath.Rel("a/b", "a/c/t/file")
	if err != nil {
		panic(err)
	}
	fmt.Println(rel)
}
