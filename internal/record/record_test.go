package record

import (
	"bytes"
	"errors"
	"reflect"
	"testing"
)

func TestEncodeProducesV1Bytes(t *testing.T) {
	// Testcase data
	checksum := []byte{0x5d, 0x6e, 0xbe, 0x22}
	timestamp := []byte{0x01, 0x02, 0x03, 0x04, 0x05, 0x06, 0x07, 0x08}
	keyLength := []byte{0x00, 0x00, 0x00, 0x02}
	valueLength := []byte{0x00, 0x00, 0x00, 0x02}
	key := []byte{0x67, 0x6f}   // hex-values for "go"
	value := []byte{0x44, 0x42} // hex-values for "DB"

	expected := make([]byte, 0, 20+len(key)+len(value))
	expected = append(expected, checksum...)
	expected = append(expected, timestamp...)
	expected = append(expected, keyLength...)
	expected = append(expected, valueLength...)
	expected = append(expected, key...)
	expected = append(expected, value...)

	rawRecord := Record{
		Timestamp: int64(0x0102030405060708),
		Key:       []byte("go"),
		Value:     []byte("DB"),
	}

	got, err := Encode(rawRecord)
	if err != nil {
		t.Fatalf("Encoding returned an error: %v", err)
	}

	if !bytes.Equal(got, expected) {
		t.Fatalf("Encode() bytes = %x, want %x", got, expected)
	}
}

func TestDecodeReadsV1Bytes(t *testing.T) {
	// Testcase data
	checksum := []byte{0x5d, 0x6e, 0xbe, 0x22}
	timestamp := []byte{0x01, 0x02, 0x03, 0x04, 0x05, 0x06, 0x07, 0x08}
	keyLength := []byte{0x00, 0x00, 0x00, 0x02}
	valueLength := []byte{0x00, 0x00, 0x00, 0x02}
	key := []byte{0x67, 0x6f}   // hex-values for "go"
	value := []byte{0x44, 0x42} // hex-values for "DB"

	data := make([]byte, 0, 20+len(key)+len(value))
	data = append(data, checksum...)
	data = append(data, timestamp...)
	data = append(data, keyLength...)
	data = append(data, valueLength...)
	data = append(data, key...)
	data = append(data, value...)

	expectedRecord := Record{
		Timestamp: int64(0x0102030405060708),
		Key:       []byte("go"),
		Value:     []byte("DB"),
	}

	got, err := Decode(data)
	if err != nil {
		t.Fatalf("Decoder returned an error: %v", err)
	}

	if !(got.Timestamp == expectedRecord.Timestamp && bytes.Equal(got.Key, expectedRecord.Key) && bytes.Equal(got.Value, expectedRecord.Value)) {
		t.Fatalf("Decoder failed.\nReturned = %+v\nExpected = %+v", got, expectedRecord)
	}
}

func TestDecodeErrIncompleteRecord(t *testing.T) {
	// Testcase data
	checksum := []byte{0x5d, 0x6e, 0xbe, 0x22}
	timestamp := []byte{0x01, 0x02, 0x03, 0x04, 0x05, 0x06, 0x07, 0x08}
	keyLength := []byte{0x00, 0x00, 0x00, 0x02}
	valueLength := []byte{0x00, 0x00, 0x00, 0x02}
	key := []byte{0x67, 0x6f} // hex-values for "go"
	value := []byte{0x44}     // hex-values for "DB" (INCOMPLETE RECORD, missing byte)

	data := make([]byte, 0, 20+len(key)+len(value))
	data = append(data, checksum...)
	data = append(data, timestamp...)
	data = append(data, keyLength...)
	data = append(data, valueLength...)
	data = append(data, key...)
	data = append(data, value...)

	record, err := Decode(data)
	if reflect.ValueOf(record).IsZero() {
		t.Fatalf("Decoder is not returning an empty instance of a record")
	}

	if !errors.Is(err, ErrIncompleteRecord) {
		t.Fatalf("Decoder returned incorrect error, expected = %v, returned = %v", ErrIncompleteRecord, err)
	}
}

func TestDecodeErrChecksumMismatch(t *testing.T) {
	// Testcase data
	checksum := []byte{0x5d, 0x6e, 0xbe, 0x22}
	timestamp := []byte{0x01, 0x02, 0x03, 0x04, 0x05, 0x06, 0x07, 0x08}
	keyLength := []byte{0x00, 0x00, 0x00, 0x02}
	valueLength := []byte{0x00, 0x00, 0x00, 0x02}
	key := []byte{0x67, 0x6f}   // hex-values for "go"
	value := []byte{0x44, 0x00} // hex-values for "DB" (SHOULD CAUSE CHECKSUM MISMATCH)

	data := make([]byte, 0, 20+len(key)+len(value))
	data = append(data, checksum...)
	data = append(data, timestamp...)
	data = append(data, keyLength...)
	data = append(data, valueLength...)
	data = append(data, key...)
	data = append(data, value...)

	record, err := Decode(data)
	if reflect.ValueOf(record).IsZero() {
		t.Fatalf("Decoder is not returning an empty instance of a record")
	}

	if !errors.Is(err, ErrChecksumMismatch) {
		t.Fatalf("Decoder returned incorrect error, expected = %v, returned = %v", ErrChecksumMismatch, err)
	}
}
