package main

import (
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
)

func checkErr(e error) {
	if e != nil {
		panic(e)
	}
}

func visit(path string, d fs.DirEntry, err error) error {
	if err != nil {
		return err
	}

	fmt.Println(" ", path, d.IsDir())
	return nil
}

func main() {
	// err := os.Mkdir("subdir", 0755) // 0755 is the permission, user, group, others, 4- read, 2 -write, 1 -exec
	// checkErr(err)

	defer os.RemoveAll("subdir") // removeAll is rm -rf

	createEmptyFile := func(name string) {
		d := []byte("")
		checkErr(os.WriteFile(name, d, 0644))
	}

	createEmptyFile("subdir/file1")

	// mkdir -p - creates hierarchy of directories
	err := os.MkdirAll("subdir/parent/child", 0755)
	checkErr(err)

	createEmptyFile("subdir/parent/file2")
	createEmptyFile("subdir/parent/file3")
	createEmptyFile("subdir/parent/child/file4")

	c, err := os.ReadDir("subdir/parent")
	checkErr(err)

	fmt.Println(c)

	for _, entry := range c {
		fmt.Println(" ", entry.Name(), entry.IsDir())
	}

	err = os.Chdir("subdir/parent/child") // cd subdir/parent/child - the same
	checkErr(err)

	c, err = os.ReadDir(".")
	checkErr(err)

	for _, entry := range c {
		fmt.Println(" ", entry.Name(), entry.IsDir())
	}

	err = os.Chdir("../../..") // cd ../../.. - the same
	checkErr(err)

	fmt.Println("Visiting subdir")
	err = filepath.WalkDir("subdir", visit)
}
