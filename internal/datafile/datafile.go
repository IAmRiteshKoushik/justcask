package datafile

import (
	"fmt"
	"os"
)

type Location struct {
	FileID uint32 // A non-negative segment identity with ample room for rotations
	Offset int64  // Go file APIs use byte offsets as int64, negative offsets are invalid
	Size   uint32 // An encoded record is a little more than 4 KB + 16MB
}

type dataFileHandle interface {
	Stat() (os.FileInfo, error)
	Write([]byte) (int, error)
	Close() error
}

type DataFile struct {
	FileID uint32
	file   dataFileHandle // os.File implements the custom interface
}

func Open(path string, fileID uint32) (*DataFile, error) {
	// Open file in append only mode and return the pointer to it
	// Permissions are 0600 instead of 0644 because DB files are:
	// - readable and writeable by owner (6)
	// - inaccessible by group (0)
	// - inaccessible by others (0)
	f, err := os.OpenFile(path, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0600)
	if err != nil {
		return nil, fmt.Errorf("Open() failed to open datafile: %w", err)
	}

	return &DataFile{
		FileID: fileID,
		file:   f,
	}, nil
}

func (f *DataFile) Append(data []byte) (Location, error) {
	var loc Location

	info, err := f.file.Stat()
	if err != nil {
		return loc, fmt.Errorf("Append(file.Stat) errored: %w", err)
	}

	fsize := info.Size() // calculate the starting offset for the record
	lenBytesWritten, err := f.file.Write(data)
	if err != nil {
		return loc, fmt.Errorf("Append(file.Write) errored: %w", err)
	}

	loc.FileID = f.FileID
	loc.Offset = fsize                 // This is where the record starts from
	loc.Size = uint32(lenBytesWritten) // Number of bytes actually written

	return loc, nil
}

func (f *DataFile) Close() error {
	return f.file.Close()
}
