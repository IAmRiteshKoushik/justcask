package record

import (
	"io"
)

type oneByteReader struct {
	reader io.Reader
}

func (r oneByteReader) Read(data []byte) (int, error) {
	if len(data) > 1 {
		data = data[:1]
	}

	return r.reader.Read(data)
}
