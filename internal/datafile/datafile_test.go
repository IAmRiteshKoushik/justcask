package datafile

import (
	"bytes"
	"errors"
	"io"
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

// Mocking is used here because we need to give the test control over failure
// that are unreliable to trigger with a real disk-file. It is easy to make
// Stat() or Write() fail or simulate it's failure without actually filling up
// the disk.
func TestAppendReturnsZeroLocationWhenStatFails(t *testing.T) {
	statErr := errors.New("stat failed")
	handle := &mockFile{statErr: statErr}

	file := &DataFile{
		FileID: 7,
		file:   handle,
	}

	location, err := file.Append([]byte("record"))
	if location != (Location{}) {
		t.Fatalf("Append(location) = %+v, want zero location", location)
	}
	if !errors.Is(err, statErr) {
		t.Fatalf("Append() error = %v, want errors.Is(err, statErr)", err)
	}
	if handle.writeCalls != 0 {
		t.Fatalf("Write() calls = %d, want 0", handle.writeCalls)
	}
}

func TestAppendReturnsZeroLocationWhenWriteIsShort(t *testing.T) {
	path := filepath.Join(t.TempDir(), "fixture")
	if err := os.WriteFile(path, []byte("existing"), 0o600); err != nil {
		t.Fatalf("WriteFile() error = %v", err)
	}
	info, err := os.Stat(path)
	if err != nil {
		t.Fatalf("Stat() error = %v", err)
	}

	data := []byte("record")
	handle := &mockFile{
		info:     info,
		writeN:   len(data) - 1,
		writeErr: io.ErrShortWrite,
	}
	file := &DataFile{
		FileID: 7,
		file:   handle,
	}

	location, err := file.Append(data)
	if location != (Location{}) {
		t.Fatalf("Append() location = %+v, want zero Location", location)
	}
	if !errors.Is(err, io.ErrShortWrite) {
		t.Fatalf("Append() error = %v, want errors.Is(err, io.ErrShortWrite)", err)
	}
	if handle.writeCalls != 1 {
		t.Fatalf("Write() calls = %d, want 1", handle.writeCalls)
	}
}

func TestAppendRefusesLaterWritesAfterWriteFailure(t *testing.T) {
	path := filepath.Join(t.TempDir(), "fixture")
	if err := os.WriteFile(path, []byte("existing"), 0o600); err != nil {
		t.Fatalf("WriteFile() error = %v", err)
	}

	info, err := os.Stat(path)
	if err != nil {
		t.Fatalf("Stat() error = %v", err)
	}

	first := []byte("first record")
	writeErr := errors.New("injected write failure")
	handle := &mockFile{
		info:     info,
		writeN:   len(first) - 1,
		writeErr: writeErr,
	}
	file := &DataFile{
		FileID: 7,
		file:   handle,
	}

	firstLocation, err := file.Append(first)
	if firstLocation != (Location{}) {
		t.Fatalf("first Append() location = %+v, want zero Location", firstLocation)
	}
	if !errors.Is(err, writeErr) {
		t.Fatalf("first Append() error = %v, want errors.Is(err, writeErr)", err)
	}

	secondLocation, err := file.Append([]byte("must not be written"))
	if secondLocation != (Location{}) {
		t.Fatalf("second Append() location = %+v, want zero Location", secondLocation)
	}
	if !errors.Is(err, ErrPoisonError) {
		t.Fatalf(
			"second Append() error = %v, want errors.Is(err, ErrPoisonError)",
			err,
		)
	}
	if handle.writeCalls != 1 {
		t.Fatalf("Write() calls = %d, want 1", handle.writeCalls)
	}
}
