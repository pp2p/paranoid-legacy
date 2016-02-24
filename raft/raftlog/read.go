package raftlog

import (
	"errors"
	pb "github.com/cpssd/paranoid/proto/raft"
	"github.com/golang/protobuf/proto"
	"io/ioutil"
	"path"
	"strconv"
)

// GetLogEntry will read an entry at the given index returning
// the protobuf and an error if something went wrong
func (rl *RaftLog) GetLogEntry(index uint64) (entry *pb.LogEntry, err error) {
	rl.indexLock.Lock()
	defer rl.indexLock.Unlock()

	if index < 1 || index >= rl.currentIndex {
		return nil, errors.New("Index out of bounds")
	}

	fileIndex := ci2fi(index)
	filePath := path.Join(rl.logDir, strconv.FormatUint(fileIndex, 10))
	fileData, err := ioutil.ReadFile(filePath)
	if err != nil {
		return nil, errors.New("Failed to read logfile")
	}

	entry = &pb.LogEntry{}
	err = proto.Unmarshal(fileData, entry)
	if err != nil {
		return nil, errors.New("Failed to Unmarshal file data")
	}

	return entry, nil
}

// GeEntriesSince returns a list of entries including and after the one
// at the given index, and an error object if somethign went wrong
func (rl *RaftLog) GetEntriesSince(index uint64) (entries []*pb.Entry, err error) {
	return rl.GetLogEntries(index, 100000000)
}

func min(a, b uint64) uint64 {
	if a < b {
		return a
	}
	return b
}

func (rl *RaftLog) GetLogEntries(index, maxCount uint64) (entries []*pb.Entry, err error) {
	rl.indexLock.Lock()
	defer rl.indexLock.Unlock()

	if index < 1 || index >= rl.currentIndex {
		return nil, errors.New("Index out of bounds")
	}

	entries = make([]*pb.Entry, min(rl.currentIndex-index, maxCount))
	for i := index; i < min(rl.currentIndex, index+maxCount); i++ {
		fileIndex := ci2fi(i)
		filePath := path.Join(rl.logDir, strconv.FormatUint(fileIndex, 10))
		fileData, err := ioutil.ReadFile(filePath)
		if err != nil {
			return nil, errors.New("Failed to read logfile")
		}

		entry := &pb.LogEntry{}
		err = proto.Unmarshal(fileData, entry)
		if err != nil {
			return nil, errors.New("Failed to Unmarshal file data")
		}

		entries[i-index] = entry.Entry
	}

	return entries, nil
}
