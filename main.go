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
	//files, err := ioutil.ReadDir(pathDir)

	err := dirTree1(out, pathDir, printFiles, 0)

	if err != nil {
		log.Fatal(err)
	}

	return err
}

func dirTree1(out *os.File, pathDir string, printFiles bool, ntabs int) error {
	files, err := ioutil.ReadDir(pathDir)
	ntabs++
	if err != nil {
		log.Fatal(err)
	}
	var prevTabs = ntabs

	for _, f := range files {
		fmt.Print(prevTabs)
		for t := 0; t < prevTabs; t++ {
			if prevTabs > 2 && t > 1 {
				fmt.Print("│")
			}
			fmt.Print("\t")
		}
		var fileName string
		fileName = f.Name()
		if len(fileName) > 3 {
			fileName = fileName[0:4]
		}

		if fileName != ".git" {

			switch mode := f.Mode(); {
			case mode.IsDir():
				// do directory stuff

				fmt.Print("├───")
				fmt.Println(f.Name())
				newDir := pathDir + string(os.PathSeparator) + f.Name()

				err = dirTree1(out, newDir, printFiles, ntabs)

			case mode.IsRegular():
				// do file stuff
				if printFiles {
					fmt.Print("├───")
					fmt.Println(f.Name())
				}
			}
		}
	}
	return err
}
