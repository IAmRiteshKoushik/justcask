package datafile

import (
	"bytes"
	"os"
	"path/filepath"
	"testing"
)

func TestAppendWritesContiguousBytesAndReturnsLocations(t *testing.T) {
	const fileID uint32 = 7

	path := filepath.Join(t.TempDir(), "00000007.data")
	file, err := Open(path, fileID)
	if err != nil {
		t.Fatalf("Open() error = %v", err)
	}
	t.Cleanup(func() {
		_ = file.Close()
	}) // Runs after the test is complete, it is internally a defer function

	first := []byte{0x10, 0x20, 0x30}
	second := []byte{0x40, 0x50}

	firstLocation, err := file.Append(first)
	if err != nil {
		t.Fatalf("first Append() error = %v", err)
	}
	wantFirstLocation := Location{
		FileID: fileID,
		Offset: 0,
		Size:   uint32(len(first)),
	}
	if firstLocation != wantFirstLocation {
		t.Fatalf("first Append() location = %+v, want +%v", firstLocation, wantFirstLocation)
	}

	secondLocation, err := file.Append(second)
	if err != nil {
		t.Fatalf("second Append() error = %v", err)
	}
	wantSecondLocation := Location{
		FileID: fileID,
		Offset: int64(len(first)),
		Size:   uint32(len(second)),
	}
	if secondLocation != wantSecondLocation {
		t.Fatalf("second Append() location = %+v, want +%v", secondLocation, wantSecondLocation)
	}

	// Byte check
	gotBytes, err := os.ReadFile(path)
	if err != nil {
		t.Fatalf("ReadFile() error = %v", err)
	}
	wantBytes := append(append([]byte{}, first...), second...)
	if !bytes.Equal(gotBytes, wantBytes) {
		t.Fatalf("fiel bytes = %x, want %x", gotBytes, wantBytes)
	}
}
