package commands

import (
	"encoding/json"
	"github.com/cpssd/paranoid/pfsm/returncodes"
	"io"
	"io/ioutil"
	"os"
	"path"
)

// LinkCommand creates a link of a file.
// args[0] is the init point, args[1] is the existing file name, args[2] is the target file name
func LinkCommand(args []string) {
	Log.Info("link command called")
	if len(args) < 3 {
		Log.Fatal("Not enough arguments!")
	}

	directory := args[0]
	existingFilePath := getParanoidPath(directory, args[1])
	targetFilePath := getParanoidPath(directory, args[2])

	Log.Verbose("link : given directory = " + directory)

	getFileSystemLock(directory, exclusiveLock)
	defer unLockFileSystem(directory)

	existingFileType := getFileType(directory, existingFilePath)
	if existingFileType == typeENOENT {
		io.WriteString(os.Stdout, returncodes.GetReturnCode(returncodes.ENOENT))
		return
	}
	if existingFileType == typeDir {
		io.WriteString(os.Stdout, returncodes.GetReturnCode(returncodes.EISDIR))
		return
	}
	if existingFileType == typeSymlink {
		io.WriteString(os.Stdout, returncodes.GetReturnCode(returncodes.EIO))
		return
	}

	if getFileType(directory, targetFilePath) != typeENOENT {
		io.WriteString(os.Stdout, returncodes.GetReturnCode(returncodes.EEXIST))
		return
	}

	// getting inode and fileMode of existing file
	inodeBytes, code := getFileInode(existingFilePath)
	if code != returncodes.OK {
		io.WriteString(os.Stdout, returncodes.GetReturnCode(code))
		return
	}
	fileInfo, err := os.Stat(existingFilePath)
	if err != nil {
		Log.Fatal("error stating existing file:", err)
	}
	fileMode := fileInfo.Mode()

	// creating target file pointing to same inode
	err = ioutil.WriteFile(targetFilePath, inodeBytes, fileMode)
	if err != nil {
		Log.Fatal("error writing to names file:", err)
	}

	// getting contents of inode
	inodePath := path.Join(directory, "inodes", string(inodeBytes))
	Log.Verbose("link : reading file " + inodePath)
	inodeContents, err := ioutil.ReadFile(inodePath)
	if err != nil {
		Log.Fatal("error reading inode:", err)
	}

	nodeData := &inode{}
	err = json.Unmarshal(inodeContents, &nodeData)
	if err != nil {
		Log.Fatal("error unmarshalling inode data:", err)
	}

	// itterating count and saving
	nodeData.Count++
	Log.Verbose("link : opening file " + inodePath)
	openedFile, err := os.OpenFile(inodePath, os.O_WRONLY, 0600)
	if err != nil {
		Log.Fatal("error opening file:", err)
	}

	Log.Verbose("link : truncating file " + inodePath)
	err = openedFile.Truncate(0)
	if err != nil {
		Log.Fatal("error truncating file:", err)
	}

	newJSONData, err := json.Marshal(&nodeData)
	if err != nil {
		Log.Fatal("error marshalling json:", err)
	}

	Log.Verbose("link : writing to file " + inodePath)
	_, err = openedFile.Write(newJSONData)
	if err != nil {
		Log.Fatal("error writing to inode file:", err)
	}

	// closing file
	Log.Verbose("link : closing file " + inodePath)
	err = openedFile.Close()
	if err != nil {
		Log.Fatal("error closing file:", err)
	}

	if !Flags.Network {
		sendToServer(directory, "link", args[1:], nil)
	}
	io.WriteString(os.Stdout, returncodes.GetReturnCode(returncodes.OK))
}
