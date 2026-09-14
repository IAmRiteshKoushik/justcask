package record

import "errors"

var (
	ErrIncompleteRecord = errors.New("record length is invalid")
	ErrChecksumMismatch = errors.New("record checksum does not match with computed checksum")
)
