package main

import (
	"fmt"
	"io"
	"os"
	"sort"
)

func dirTree(writer io.Writer, path string, printFiles bool) error {
	err := dirTreeRecurs(writer, path, printFiles, 1)
	if err != nil {
		return err
	}

	return nil
}

func dirTreeRecurs(writer io.Writer, path string, printFiles bool, level int) error {
	file, err := os.Open(path)
	if err != nil {
		return err
	}

	files, err := file.Readdir(100)
	if err != nil {
		return err
	}

	sort.Slice(files, func(i, j int) bool {
		return (files[i].IsDir() && files[i].Name() < files[j].Name()) || (!files[i].IsDir() && files[i].Name() < files[j].Name())
	})

	for idx, fileInfo := range files {
		if !printFiles && !fileInfo.IsDir() {
			continue
		}
		isLastItem := (idx + 1) == len(files)
		printLevel(writer, level, isLastItem)
		printFileName(writer, fileInfo)
		if fileInfo.IsDir() {
			newPath := path + "/" + fileInfo.Name()
			level++
			dirTreeRecurs(writer, newPath, printFiles, level)
			level--
		}
	}

	return nil
}

func printFileName(writer io.Writer, fileInfo os.FileInfo) {
	if fileInfo.IsDir() {
		fmt.Fprintf(writer, "%s\n", fileInfo.Name())
	} else {
		if fileInfo.Size() != 0 {
			fmt.Fprintf(writer, "%s (%db)\n", fileInfo.Name(), fileInfo.Size())
		} else {
			fmt.Fprintf(writer, "%s (empty)\n", fileInfo.Name())
		}
	}
}

func printLevel(writer io.Writer, level int, isLastItem bool) {
	for i := 0; i < level; i++ {
		if i == (level - 1) {
			if !isLastItem {
				fmt.Fprint(writer, "├───")
			} else {
				fmt.Fprint(writer, "└───")
			}
		} else {
			if isLastItem && i == (level-2) {
				fmt.Fprintf(writer, "\t")
			} else {
				fmt.Fprint(writer, "│\t")
			}
		}
	}
}

func main() {
	out := os.Stdout
	if !(len(os.Args) == 2 || len(os.Args) == 3) {
		panic("usage go run main.go . [-f]")
	}
	path := os.Args[1]
	printFiles := len(os.Args) == 3 && os.Args[2] == "-f"
	err := dirTree(out, path, printFiles)
	if err != nil {
		panic(err.Error())
	}
}
