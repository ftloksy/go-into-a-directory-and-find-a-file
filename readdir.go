package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

// given a pair of a directory path and a filename
func readDir(dir string, findFile string) []string {
	var paths []string
	files, err := ioutil.ReadDir(dir)
	if err != nil {
		fmt.Println(err)
	}

	for _, file := range files {
		if file.IsDir() {
			directory := file.Name()
			directoryName := dir + "/" + directory
			// ... is very interesting
			paths = append(paths, readDir(directoryName, findFile)...)
		} else {
			filename := file.Name()
			fileName := dir + "/" + filename
			if findFile == file.Name() {
				fmt.Println("File: " + fileName)
				paths = append(paths, fileName)
			}
		}
	}

	// returns an unordered collection of
	// all the subdirectory paths in the directories
	// that contain files with the same filename.
	return paths
}

func main() {
	findFile := "main"
	dir, err := os.Getwd()
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("From Directory: " + dir)
	fmt.Println("Find File: " + findFile)
	readDir(dir, findFile)

}
