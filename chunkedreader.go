// Package chunkedreader is a light weight wrapper for bufio that enables
// reading of byte streams in fixed size chunks.

// Kailash Nadh, http://nadh.in/code/chunkedreader
// March 2015
// MIT License

package chunkedreader

import (
	"bufio"
	"io"
)

type ChunkedReader struct {
	reader  *bufio.Reader
	scanner *bufio.Scanner
	length  int
}

// Create a new instance of the ChunkedReader.
func New(rd io.Reader, length int) *ChunkedReader {
	r := bufio.NewReader(rd)
	c := &ChunkedReader{
		reader:  r,
		scanner: bufio.NewScanner(r),
		length:  length,
	}

	// Assign the custom split function to the bufio Scanner
	c.scanner.Split(c.split)

	return c
}

func (c *ChunkedReader) Read() bool {
	return c.scanner.Scan()
}

func (c *ChunkedReader) Bytes() []byte {
	return c.scanner.Bytes()
}

// The custom bufio.Scanner split function. While it's meant
// to split streams by delimiters, we're splitting based on
// fixed lengths here.
func (c *ChunkedReader) split(data []byte, atEOF bool) (advance int, token []byte, err error) {
	if atEOF && len(data) == 0 {
		return 0, nil, nil
	}

	// Read a fixed size.
	if len(data) > c.length {
		return c.length, data[0:c.length], nil
	}

	if atEOF {
		return len(data), data, nil
	}

	// Request more data.
	return 0, nil, nil
}
