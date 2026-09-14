package record

import "errors"

var (
	ErrIncompleteRecord = errors.New("record length is invalid")
	ErrChecksumMismatch = errors.New("record checksum does not match with computed checksum")
	ErrRecordTooLarge   = errors.New("record exceeds size limits of 4KB key, 16MB value")
)
