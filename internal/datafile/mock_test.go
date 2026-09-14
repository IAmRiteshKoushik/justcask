package datafile

import "os"

type mockFile struct {
	info    os.FileInfo
	statErr error

	writeN     int
	writeErr   error
	writeCalls int
}

func (f *mockFile) Stat() (os.FileInfo, error) {
	return f.info, f.statErr
}

func (f *mockFile) Write([]byte) (int, error) {
	f.writeCalls++
	return f.writeN, f.writeErr
}

func (f *mockFile) Close() error {
	return nil
}
