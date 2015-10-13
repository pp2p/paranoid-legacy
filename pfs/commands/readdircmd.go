package commands

import (
	"fmt"
	"io/ioutil"
	"log"
	"path"
)

func ReadDirCommand(args []string) {
	if len(args) < 1 {
		log.Fatal("Not enough arguments!")
	}
	directory := args[0]
	files, err := ioutil.ReadDir(path.Join(directory, "/names/"))
	checkErr("readdir", err)
	for i := 0; i < len(files); i++ {
		fmt.Println(files[i].Name())
	}
}