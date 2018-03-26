package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

func main() {
	out := os.Stdout
	if !(len(os.Args) == 2 || len(os.Args) == 3) {
		panic("usage go run main.go . [-f]")
	}
	pathDir := os.Args[1]
	printFiles := len(os.Args) == 3 && os.Args[2] == "-f"

	err := dirTree(out, pathDir, printFiles)

	if err != nil {
		panic(err.Error())
	}
}

func dirTree(out *os.File, pathDir string, printFiles bool) error {
	files, err := ioutil.ReadDir(pathDir)

	// sort.Strings(files)

	if err != nil {
		log.Fatal(err)
	}

	for _, f := range files {
		switch mode := f.Mode(); {
		case mode.IsDir():
			// do directory stuff
			fmt.Println("dir: ", f.Name())
			newDir := pathDir + string(os.PathSeparator) + f.Name()
			err = dirTree(out, newDir, printFiles)

		case mode.IsRegular():
			// do file stuff
			if printFiles {
				fmt.Println("file: ", f.Name())
			}
		}
	}
	return err
}
