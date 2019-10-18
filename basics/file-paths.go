package main

// filepath package provides functions to parse and construct file paths
// that s portable between os

import (
	"fmt"
	"path/filepath"
	"strings"
)

func main() {

	// usse join to construct paths portably
	p := filepath.Join("dir1", "dir2", "filename")
	fmt.Println("p:", p)

	fmt.Println(filepath.Join("dir1//", "filename"))
	fmt.Println(filepath.Join("dir1/../dir1", "filename"))

	// Dir and base can be used to split a path to dir and the file
	// Split will do the same thing

	fmt.Println("Dir(p):", filepath.Dir(p))
	fmt.Println("Base(p)", filepath.Base(p))

	// Check if a path is absolute
	fmt.Println(filepath.IsAbs("dir/file"))
	fmt.Println(filepath.IsAbs("/dir/file"))

	filename := "config.json"

	//some file names have extensions following a dot. Split the extension out of such names with Ext
	ext := filepath.Ext(filename)
	fmt.Println(ext)

	// find file's name with extension removed use strings.TrimSuffix
	fmt.Println(strings.TrimSuffix(filename, ext))

	// Rel finds a relative path between a ase and target . Error if target cannot be made relative to base
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
