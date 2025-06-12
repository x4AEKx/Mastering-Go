package main

import (
	"fmt"
	"os"
	"path/filepath"
)

func main() {
	args := os.Args
	if len(args) == 1 {
		fmt.Println("Please provide some arguments!")
		return
	}

	files := args[1:]

	path := os.Getenv("PATH")
	pathSplit := filepath.SplitList(path)

	for _, file := range files {
		for _, directory := range pathSplit {
			fullPath := filepath.Join(directory, file)
			fileInfo, err := os.Stat(fullPath)
			if err == nil {
				mode := fileInfo.Mode()

				if mode.IsRegular() {
					if mode&0111 != 0 {
						fmt.Println(fullPath)
					}
				}
			}
		}
	}

}
