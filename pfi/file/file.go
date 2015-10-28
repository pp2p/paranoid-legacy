package file

import (
	"github.com/cpssd/paranoid/pfi/pfsinterface"
	"github.com/cpssd/paranoid/pfi/util"
	"strconv"

	"github.com/hanwen/go-fuse/fuse"
	"github.com/hanwen/go-fuse/fuse/nodefs"
)

var mountPoint string

//ParanoidFile is a custom file struct with read and write functions
type ParanoidFile struct {
	Name string
	nodefs.File
}

//NewParanoidFile returns a new object of ParanoidFile
func NewParanoidFile(name string) nodefs.File {
	return &ParanoidFile{
		Name: name,
		File: nodefs.NewDefaultFile(),
	}
}

//Read reads a file and returns an array of bytes
func (f *ParanoidFile) Read(buf []byte, off int64) (fuse.ReadResult, fuse.Status) {
	util.LogMessage("Read called on : " + f.Name)
	data := pfsinterface.RunCommand(nil, "read", util.PfsInitPoint, f.Name, strconv.FormatInt(off, 10), strconv.FormatInt(int64(len(buf)), 10))
	return fuse.ReadResultData(data), fuse.OK
}

//Write writes to a file
func (f *ParanoidFile) Write(content []byte, off int64) (uint32, fuse.Status) {
	util.LogMessage("Write called on : " + f.Name)
	pfsinterface.RunCommand(content, "write", util.PfsInitPoint, f.Name, strconv.FormatInt(off, 10), strconv.FormatInt(int64(len(content)), 10))
	return uint32(len(content)), fuse.OK
}