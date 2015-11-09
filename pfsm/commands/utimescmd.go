package commands

import (
	"encoding/json"
	"github.com/cpssd/paranoid/pfsm/returncodes"
	"io"
	"io/ioutil"
	"log"
	"os"
	"path"
	"syscall"
	"time"
)

type timeInfo struct {
	Atime *time.Time `json:"atime",omitempty`
	Mtime *time.Time `json:"mtime",omitempty`
}

//UtimesCommand updates the acess time and modified time of a file
func UtimesCommand(args []string) {
	verboseLog("utimes command called")
	if len(args) < 2 {
		log.Fatalln("Not enough arguments!")
	}
	directory := args[0]
	verboseLog("utimes : given directory = " + directory)

	if !checkFileExists(path.Join(directory, "names", args[1])) {
		io.WriteString(os.Stdout, returncodes.GetReturnCode(returncodes.ENOENT))
		return
	}

	input, err := ioutil.ReadAll(os.Stdin)
	checkErr("utimes", err)
	times := timeInfo{}
	err = json.Unmarshal(input, &times)
	checkErr("utimes", err)

	fileNameBytes, err := ioutil.ReadFile(path.Join(directory, "names", args[1]))
	checkErr("utimes", err)
	fileName := string(fileNameBytes)

	err = syscall.Access(path.Join(directory, "contents", fileName), getAccessMode(syscall.O_WRONLY))
	if err != nil {
		io.WriteString(os.Stdout, returncodes.GetReturnCode(returncodes.EACCES))
		return
	}

	file, err := os.Open(path.Join(directory, "contents", fileName))
	checkErr("utimes", err)

	fi, err := file.Stat()
	stat := fi.Sys().(*syscall.Stat_t)
	oldatime := time.Unix(int64(stat.Atim.Sec), int64(stat.Atim.Nsec))
	oldmtime := fi.ModTime()
	if times.Atime == nil && times.Mtime == nil {
		log.Fatalln("utimes : no times to update!")
	}

	if times.Atime == nil {
		os.Chtimes(path.Join(directory, "contents", fileName), oldatime, *times.Mtime)
	} else if times.Mtime == nil {
		os.Chtimes(path.Join(directory, "contents", fileName), *times.Atime, oldmtime)
	} else {
		os.Chtimes(path.Join(directory, "contents", fileName), *times.Atime, *times.Mtime)
	}

	io.WriteString(os.Stdout, returncodes.GetReturnCode(returncodes.OK))
}