package record

import (
	"bytes"
	"errors"
	"io"
	"reflect"
	"testing"
)

func TestReadFromEmptyReaderReturnsEOF(t *testing.T) {
	got, size, err := ReadFrom(bytes.NewReader(nil))

	if !reflect.ValueOf(got).IsZero() {
		t.Fatalf("ReadFrom() record = %+v, want zero Record", got)
	}
	if size != 0 {
		t.Fatalf("ReadFrom() size = %d, want 0", size)
	}
	if !errors.Is(err, io.EOF) {
		t.Fatalf("ReadFrom() error = %v, want errors.Is(err, io.EOF)", err)
	}
}

func TestReadFromReturnsUnexpectedEOFForEveryNonEmptyTruncatedPrefix(t *testing.T) {
	valid := validV1Bytes()

	for length := range valid {
		if length == 0 {
			continue
		}

		got, _, err := ReadFrom(bytes.NewReader(valid[:length]))
		if !reflect.ValueOf(got).IsZero() {
			t.Fatalf(
				"ReadFrom(%d-byte prefix) record = %+v, want zero Record",
				length,
				got,
			)
		}
		if !errors.Is(err, io.ErrUnexpectedEOF) {
			t.Fatalf(
				"ReadFrom(%d-byte prefix) error = %v, want errors.Is(err, io.ErrUnexpectedEOF)",
				length,
				err,
			)
		}
	}
}

func TestReadFromReturnsChecksumMismatchForCompleteCorruptRecord(t *testing.T) {
	data := append([]byte(nil), validV1Bytes()...)
	data[len(data)-1] ^= 0x01 // XOR is to cause a byte flip and trigger checksum mismatch

	got, _, err := ReadFrom(bytes.NewReader(data))
	if !reflect.ValueOf(got).IsZero() {
		t.Fatalf("ReadFrom() record = %+v, want zero Record", got)
	}
	if !errors.Is(err, ErrChecksumMismatch) {
		t.Fatalf(
			"ReadFrom() error = %v, want errors.Is(err, ErrChecksumMismatch)",
			err,
		)
	}
}
