package record

import (
	"encoding/binary"
	"hash/crc32"
)

type Record struct {
	Timestamp int64
	Key       []byte
	Value     []byte
}

const headerLength = 20

func Encode(record Record) ([]byte, error) {
	offset := headerLength
	data := make([]byte, offset+len(record.Key)+len(record.Value))

	binary.BigEndian.PutUint64(data[4:12], uint64(record.Timestamp))
	binary.BigEndian.PutUint32(data[12:16], uint32(len(record.Key)))
	binary.BigEndian.PutUint32(data[16:20], uint32(len(record.Value)))

	copy(data[offset:], record.Key)
	offset += len(record.Key)
	copy(data[offset:], record.Value)

	checksum := crc32.ChecksumIEEE(data[4:])
	binary.BigEndian.PutUint32(data[0:4], checksum)

	return data, nil
}

func Decode(data []byte) (Record, error) {
	var record Record
	recordLength := len(data)

	// Rejecting data shorter than 20 bytes
	if recordLength < headerLength {
		return record, ErrIncompleteRecord
	}

	// Metadata fields
	checksum := binary.BigEndian.Uint32(data[0:4])
	timestamp := int64(binary.BigEndian.Uint64(data[4:12]))
	kLength := int(binary.BigEndian.Uint32(data[12:16]))
	vLength := int(binary.BigEndian.Uint32(data[16:20]))

	// Verify declared length
	declaredLength := 20 + kLength + vLength
	if recordLength != declaredLength {
		return record, ErrIncompleteRecord
	}

	// Verify checksum
	computedChecksum := crc32.ChecksumIEEE(data[4:])
	if computedChecksum != checksum {
		return record, ErrChecksumMismatch
	}

	record.Timestamp = timestamp
	record.Key = data[20 : 20+kLength]
	record.Value = data[20+kLength : 20+kLength+vLength]

	return record, nil
}
