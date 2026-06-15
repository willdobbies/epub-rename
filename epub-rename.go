package main

import (
	"fmt"
	"os"
	"path/filepath"

	epub "github.com/ArcadiaLin/go-epub"
)

func rename(path string) bool {
	head, _ := filepath.Split(path)

	book, err := epub.ReadBook(path)
	if err != nil {
		fmt.Printf("Error loading %s: %s\n", path, err)
		return false
	}

	title, err := book.Title()
	if err != nil {
		fmt.Printf("Error reading title: %s\n", err)
		return false
	}

	creator, err := book.Creator()
	if err != nil {
		fmt.Printf("Error reading creator: %s\n", err)
		return false
	}

	var outname = fmt.Sprintf("%s - %s.epub", creator, title)
	var outpath = filepath.Join(head, outname)

	err = os.Rename(path, outpath)
	if err != nil {
		fmt.Printf("Error reading creator: %s\n", err)
		return false
	}

	fmt.Printf("Renamed %s -> %s\n", path, outpath)
	return true
}

func main() {
	var path = os.Args[1]
	rename(path)
}
