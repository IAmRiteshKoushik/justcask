package record

import "io"

// Recovery should ask the record package for one validated reocrd, not
// duplicate header parsing in the datafile layer too.
func ReadFrom(r io.Reader) (Record, int, error) {
	panic("TODO")
}
