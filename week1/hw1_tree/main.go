package main

import (
	"fmt"
	"io"
	"os"
)

func DirList(out io.Writer, path string, isFile bool, level string) {
	curDir, error := os.Open(path)
	if error != nil {
		fmt.Printf("err= %s\n", error)
	}
	curDirList, _ := curDir.Readdir(0)
	curDir.Close()
	if !isFile {
		tmp := curDirList[:0]
		for _, val := range curDirList {
			if val.IsDir() {
				tmp = append(tmp, val)
			}
		}
		curDirList = make([]os.FileInfo, len(tmp), len(tmp))
		copy(curDirList, tmp)
	}

	listLen := len(curDirList)
	for num, val := range curDirList {
		var levelf, slevel string
		if num >= 0 {
			levelf = "├"
			slevel = "│"
		}
		if num == listLen-1 {
			levelf = "└"
			slevel = ""
		}

		if val.IsDir() {
			// level++
			fmt.Fprintf(out, "%s%s───%s\n", level, levelf, val.Name())
			DirList(out, path+string(os.PathSeparator)+val.Name(), isFile, level+slevel+"	")
		} else {
			if val.Size() == 0 {
				fmt.Fprintf(out, "%s%s───%s (empty)\n", level, levelf, val.Name())
			} else {
				fmt.Fprintf(out, "%s%s───%s (%db)\n", level, levelf, val.Name(), val.Size())

			}
		}
	}
}

func dirTree(out io.Writer, path string, printFiles bool) error {
	DirList(out, path, printFiles, "")
	return nil
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
	// path, _ := os.Getwd()
	// DirList(path, true, "")
	// d := os.

}
